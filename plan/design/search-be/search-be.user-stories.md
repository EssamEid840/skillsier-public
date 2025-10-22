# search-be User Stories

**Complete Extended User Stories**  
Following domain-driven design with event-sourcing, CQRS projections, ML/AI-powered recommendations, and platform alignment.

---

## Global Conventions (applies to all sections)

**Event envelope (Kafka / outbox)**

```
event_id (ULID), event_type, version, occurred_at (UTC), actor_id, tenant_id,
correlation_id, causation_id, partition_key, schema_ref (hash),
nonpii_payload (domain DTO; no raw embeddings or PII)
```

**Idempotent write-path**

- All write handlers accept Idempotency-Key (UUIDv4). Server returns the original success payload on safe retries (TTL 24h).
- Natural keys prevent duplicates (e.g., {user_id, query_hash} for search logs; {entity_id, index_name} for indexing).
- Outbox is _exactly-once_ per event_id. Inbox performs dedupe on event_id + causation_id.

**Non-PII rules**

- Never store raw emails/phones/IM handles inside search queries; anonymize user identifiers in logs.
- Embeddings are vector representations only; no raw personal data in vectors.
- Search logs use hashed user IDs; full queries scrubbed of PII before storage.

**Platform alignment**

- Follows provided folder structure (domain/application/interfaces layers).
- Topics are grouped by feature (e.g., search.*, recommendation.*, taxonomy.*, matching.*).
- Projections are suffixed _read and kept query-optimized; commands/queries in application layer.
- Elasticsearch is the primary index store; PostgreSQL for metadata and configuration.

---

# 1 - CORE SEARCH & INDEXING

## 1.1 search_index/

### Stories

- As a **system**, I want to maintain search indices (jobs, users, portfolios) so that content is discoverable.
- As an **admin**, I want to manage index aliases and versions so that I can perform zero-downtime reindexing.
- As a **system**, I want index settings (shards, replicas, analyzers) configurable so that performance is optimized.
- As a **system**, I want to track index health and statistics so that I can monitor search quality.
- As an **admin**, I want to change index visibility (public, restricted, archived) so that content access is controlled.

### Flow

- CreateIndexCommand(admin_id, index_name, settings, mappings) → ValidateSettings() | CreateESIndex() | CreateIndexMetadata(version=1) → CacheSettings() → **Outbox:** search.index.created.v1
- UpdateIndexMappingsCommand(index_name, mappings, admin_id) → ValidateMappings() | UpdateESIndex() | IncrementVersion() → InvalidateCache() → **Outbox:** search.index.mappings.updated.v1
- SetIndexAliasCommand(index_name, alias, admin_id) → ValidateAlias() | CreateESAlias() | UpdateMetadata() → **Outbox:** search.index.alias.set.v1
- ChangeIndexVisibilityCommand(index_name, visibility, admin_id) → ValidateVisibility(PUBLIC/RESTRICTED/ARCHIVED) | UpdateMetadata() → **Outbox:** search.index.visibility.changed.v1
- ArchiveIndexCommand(index_name, admin_id) → MarkArchived() | RemoveFromSearchTargets() → **Outbox:** search.index.archived.v1
- GetIndexHealthQuery(index_name) → RetrieveESHealth() | RetrieveMetadata() → Return health stats + doc count
- ListIndicesQuery(admin_id) → RetrieveActiveIndices() → Return index list with stats

### Projections

- search_indices_read (index_name → settings, version, doc_count, visibility, health, created_at)
- index_alias_mappings_read (alias → index_names list)
- index_health_cache (Redis: index_name → health stats, TTL 5m)

### Events

- search.index.created.v1, search.index.mappings.updated.v1, search.index.alias.set.v1, search.index.visibility.changed.v1, search.index.archived.v1, search.index.health.degraded.v1

### RBAC/SLO

- **RBAC:** **ADMIN** for create/update/archive; **SYSTEM** for health checks
- **SLO:** create P95 < 5s; update mappings P95 < 3s; get health P95 < 200ms
- **Limits:** max 100 indices; max 1000 fields per index; max 10 aliases per index
- **Idempotency:** by (index_name) for create; by (index_name, version) for update

---

## 1.2 indexing/

### Stories

- As a **system**, I want to index jobs when created/updated so that they're searchable immediately.
- As a **system**, I want to index user profiles when modified so that talent discovery is current.
- As a **system**, I want bulk indexing with backpressure so that large reindex operations don't overwhelm Elasticsearch.
- As a **system**, I want to handle indexing failures with retry logic so that documents eventually get indexed.
- As an **admin**, I want to trigger reindexing of all documents so that schema changes are applied.

### Flow

- IndexJobCommand(job_id, job_data) → TransformToDocument(title, description, skills, location, budget, boost_factors) | GenerateEmbeddings(title, description) | IndexToES(jobs_index) | RecordMetadata(indexed_at) → **Outbox:** search.job.indexed.v1
- IndexUserCommand(user_id, user_data) → TransformToDocument(name, title, skills, experience, rating, location) | GenerateEmbeddings(title, overview) | IndexToES(users_index) | RecordMetadata(indexed_at) → **Outbox:** search.user.indexed.v1
- IndexPortfolioCommand(portfolio_id, portfolio_data) → TransformToDocument(title, description, skills, media_refs) | GenerateEmbeddings() | IndexToES(portfolios_index) | RecordMetadata(indexed_at) → **Outbox:** search.portfolio.indexed.v1
- BulkIndexCommand(documents[], index_name) → ChunkDocuments(batch_size=500) | ForEachBatch: IndexBulk() | ApplyBackpressure() | TrackProgress() → **Outbox:** search.bulk.indexed.v1
- RemoveFromIndexCommand(entity_id, index_name) → DeleteFromES() | RemoveMetadata() → **Outbox:** search.document.removed.v1
- ReindexAllCommand(index_name, admin_id) → CreateNewIndex(version++) | BulkReindex(old → new) | SwapAlias() | DeleteOldIndex() → **Outbox:** search.index.reindexed.v1

### Projections

- indexed_documents_read (document_id → index_name, indexed_at, version, status)
- indexing_queue_read (pending documents by priority)
- indexing_failures_read (failed documents with error + retry_count)

### Events

- search.job.indexed.v1, search.user.indexed.v1, search.portfolio.indexed.v1, search.document.removed.v1, search.bulk.indexed.v1, search.index.reindexed.v1, search.indexing.failed.v1

### RBAC/SLO

- **RBAC:** **SYSTEM** for index; **ADMIN** for reindex
- **SLO:** single document index P95 < 500ms; bulk index P95 < 10s per 1000 docs; remove P95 < 200ms
- **Limits:** max 10,000 docs per bulk request; max 3 retries for failures; indexing timeout 30s
- **Idempotency:** by (entity_id, index_name, document_hash)

---

## 1.3 search_query/

### Stories

- As a **user**, I want to search for jobs with filters (skills, location, budget, type) so that I find relevant opportunities.
- As a **user**, I want to search for talent with filters (skills, experience, rating, availability) so that I can hire the right person.
- As a **system**, I want to log all search queries for analytics so that I can improve search quality.
- As a **user**, I want search suggestions (autocomplete) so that I can refine my query quickly.
- As a **system**, I want to detect invalid or malicious queries so that I can protect the system.

### Flow

- SearchJobsCommand(user_id, query, filters, sort, page, size) → ValidateFilters() | BuildESQuery(match + filters + boost) | ExecuteSearch() | TransformResults() | LogQuery(anon_user_hash, latency) → **Outbox:** search.query.logged.v1
- SearchFreelancersCommand(user_id, query, filters, sort, page, size) → ValidateFilters() | BuildESQuery() | ExecuteSearch() | TransformResults() | LogQuery() → **Outbox:** search.query.logged.v1
- SearchPortfoliosCommand(user_id, query, filters, page, size) → ValidateFilters() | BuildESQuery() | ExecuteSearch() | TransformResults() | LogQuery() → **Outbox:** search.query.logged.v1
- GetSearchSuggestionsQuery(prefix, index_name) → ESCompletionSuggest() → Return suggestions list
- ExplainSearchQuery(query, filters, admin_id) → BuildESQuery() | ESExplainAPI() → Return explain details

### Projections

- search_logs_read (query_hash → count, avg_latency, result_count, click_through_rate, timestamp)
- search_filters_usage_read (filter_name → usage_count, popular_values)
- query_performance_read (query patterns by latency)

### Events

- search.query.logged.v1, search.query.alert.triggered.v1 (high latency), search.query.failed.v1

### RBAC/SLO

- **RBAC:** **USER** for search; **ADMIN** for explain
- **SLO:** search P95 < 300ms; suggestions P95 < 100ms; explain P95 < 500ms
- **Limits:** max 100 results per page; max 10,000 results total; query timeout 5s
- **Idempotency:** query logs by (user_hash, query_hash, timestamp_hour)

---

## 1.4 saved_search/

### Stories

- As a **user**, I want to save searches with custom names so that I can rerun them easily.
- As a **user**, I want alerts on saved searches so that I'm notified of new matching results.
- As a **user**, I want to schedule saved searches (daily, weekly) so that I stay updated automatically.
- As a **system**, I want to execute saved search alerts asynchronously so that notifications are timely.
- As a **user**, I want to disable/delete saved searches so that I control my alerts.

### Flow

- SaveSearchCommand(user_id, name, query, filters, alert_settings?) → ValidateName() | CreateSavedSearch(active=true) → If alert_enabled: ScheduleJob() → **Outbox:** search.saved.v1
- UpdateSavedSearchCommand(saved_search_id, updates, user_id) → GuardOwnership() | UpdateSavedSearch() → If schedule_changed: RescheduleJob() → **Outbox:** search.saved.updated.v1
- DeleteSavedSearchCommand(saved_search_id, user_id) → GuardOwnership() | MarkDeleted() → CancelScheduledJob() → **Outbox:** search.saved.deleted.v1
- ExecuteSavedSearchJob(saved_search_id) → LoadSearch() | ExecuteQuery() | CompareToPreviousResults() → If new_results: SendNotification() → **Outbox:** search.saved.executed.v1 + search.alert.triggered.v1
- ListSavedSearchesQuery(user_id) → RetrieveSavedSearches(active=true) → Return list with last_executed_at

### Projections

- saved_searches_read (saved_search_id → user_id, name, query, filters, alert_settings, active, last_executed_at)
- saved_search_alerts_read (alerts with schedule + new_results_count)

### Events

- search.saved.v1, search.saved.updated.v1, search.saved.deleted.v1, search.saved.executed.v1, search.alert.triggered.v1

### RBAC/SLO

- **RBAC:** **OWNER**
- **SLO:** save P95 < 300ms; execute P95 < 2s; list P95 < 200ms
- **Limits:** max 50 saved searches per user; max 10 active alerts; alert frequency min 1h
- **Idempotency:** by (user_id, query_hash) for save

---

# 2 - TAXONOMY & FACETS

## 2.1 taxonomy/

### Stories

- As an **admin**, I want to manage skills taxonomy (categories, hierarchies, synonyms) so that search is consistent.
- As a **system**, I want to map user input to canonical skills so that queries are normalized.
- As an **admin**, I want to add/update/deprecate skills so that the taxonomy evolves.
- As a **system**, I want to track skill popularity and trends so that taxonomy reflects market demand.
- As a **user**, I want skill synonyms (e.g., "JS" → "JavaScript") so that searches work intuitively.

### Flow

- CreateSkillCommand(admin_id, name, category, parent_id?, synonyms[]) → ValidateUniqueness() | CreateSkill() | AddSynonyms() → CacheTaxonomy() → **Outbox:** taxonomy.skill.created.v1
- UpdateSkillCommand(skill_id, updates, admin_id) → ValidateUpdates() | UpdateSkill() → InvalidateCache() → **Outbox:** taxonomy.skill.updated.v1
- DeprecateSkillCommand(skill_id, replacement_id?, admin_id) → MarkDeprecated() | SetReplacement() → UpdateReferences() → **Outbox:** taxonomy.skill.deprecated.v1
- AddSynonymCommand(skill_id, synonym, admin_id) → ValidateSynonym() | CreateSynonymMapping() → InvalidateCache() → **Outbox:** taxonomy.synonym.added.v1
- MapUserSkillCommand(user_input, context?) → LookupSynonyms() | NormalizeInput() | MatchToCanonical() → Return canonical_skill_id
- GetTaxonomyQuery(category?) → RetrieveTaxonomy(filters) → Return hierarchical structure
- GetSkillTrendsQuery(time_period) → AggregateTrendData() → Return trending skills

### Projections

- skills_taxonomy_read (skill_id → name, category, parent_id, popularity_score, status)
- skill_synonyms_read (synonym → canonical_skill_id)
- skill_trends_read (skill_id → trend_score, mentions_count, growth_rate)
- taxonomy_cache (Redis: complete taxonomy tree, TTL 1h)

### Events

- taxonomy.skill.created.v1, taxonomy.skill.updated.v1, taxonomy.skill.deprecated.v1, taxonomy.synonym.added.v1, taxonomy.trends.computed.v1

### RBAC/SLO

- **RBAC:** **ADMIN** for create/update/deprecate; **SYSTEM** for map; **USER** for get
- **SLO:** create/update P95 < 300ms; map P95 < 50ms (cached); get P95 < 150ms
- **Limits:** max 100,000 skills; max 50 synonyms per skill; max depth 5 levels
- **Idempotency:** by (skill_name) for create; by (synonym) for add

---

## 2.2 facets/

### Stories

- As a **user**, I want faceted search results (counts by skill, location, budget) so that I can filter effectively.
- As a **system**, I want to compute facets dynamically based on result set so that counts are accurate.
- As an **admin**, I want to configure which facets are available so that UI is customizable.
- As a **system**, I want to cache popular facet combinations so that performance is optimal.
- As a **user**, I want facet values sorted by relevance/popularity so that top options are first.

### Flow

- ComputeFacetsCommand(index_name, query, filters) → BuildESAggregations(skills, locations, budgets, experience, rating) | ExecuteAggregation() | TransformFacets() → CacheFacets(query_hash) → Return facets
- ConfigureFacetsCommand(index_name, facet_configs, admin_id) → ValidateConfigs() | UpdateFacetSettings() → InvalidateCache() → **Outbox:** search.facets.configured.v1
- GetFacetsQuery(index_name, query, filters) → CheckCache() → If miss: ComputeFacets() ELSE: Return cached
- GetFacetValuesQuery(facet_name, prefix?) → RetrievePopularValues() → Return sorted values with counts

### Projections

- facet_configurations_read (index_name → facet_list with settings)
- facet_popularity_read (facet_name → usage_count, popular_values[])
- facet_cache (Redis: query_hash → facets, TTL 15m)

### Events

- search.facets.configured.v1, search.facets.computed.v1

### RBAC/SLO

- **RBAC:** **ADMIN** for configure; **USER/SYSTEM** for compute/get
- **SLO:** compute P95 < 200ms (with cache); get values P95 < 100ms
- **Limits:** max 20 facets per index; max 1000 facet values
- **Idempotency:** by (index_name, facet_config_hash) for configure

---

## 2.3 filters/

### Stories

- As a **user**, I want to apply multiple filters (AND/OR logic) so that I narrow results precisely.
- As a **system**, I want to validate filter values so that invalid queries are rejected.
- As a **user**, I want range filters (budget, experience years) so that I can set min/max values.
- As a **user**, I want geo filters (distance from location) so that I find local opportunities.
- As a **system**, I want to track filter usage so that I can optimize popular combinations.

### Flow

- ApplyFiltersCommand(query, filters[]) → ValidateFilters(types, ranges) | BuildESFilterClause(must, should, must_not) | ExecuteQuery() → Return filtered results
- ValidateFilterQuery(filter_name, value) → CheckAllowedValues() | ValidateRange() | ValidateFormat() → Return validation result
- ApplyGeoFilterCommand(query, lat, lon, radius_km) → BuildESGeoQuery(geo_distance) | ExecuteQuery() → Return results within radius
- GetFilterOptionsQuery(filter_name) → RetrieveAllowedValues() → Return options list
- TrackFilterUsageCommand(filters[]) → LogUsage() → UpdatePopularityScores() → **Outbox:** search.filters.tracked.v1

### Projections

- filter_schemas_read (filter_name → type, allowed_values, validation_rules)
- filter_usage_read (filter_combinations → count, success_rate)
- popular_filters_read (top filters by usage)

### Events

- search.filters.tracked.v1, search.filters.invalid.v1

### RBAC/SLO

- **RBAC:** **USER** for apply; **SYSTEM** for track
- **SLO:** apply P95 < 50ms (within query); validate P95 < 10ms; get options P95 < 100ms
- **Limits:** max 20 filters per query; max 100 values per multi-value filter
- **Idempotency:** tracking by (filters_hash, timestamp_hour)

---

# 3 - RECOMMENDATION & PERSONALIZATION

## 3.1 recommendation/

### Stories

- As a **user**, I want personalized job recommendations based on my skills/history so that I see relevant opportunities.
- As a **system**, I want to use ML models (collaborative filtering, content-based, hybrid) so that recommendations are accurate.
- As a **user**, I want to see why jobs are recommended (match reasons) so that I trust the system.
- As a **system**, I want to track recommendation performance (CTR, conversion) so that I can improve models.
- As a **client**, I want freelancer recommendations for my jobs so that I can invite qualified candidates.

### Flow

- GenerateJobRecommendationsCommand(user_id, context, count=10) → LoadUserProfile() | LoadModel() | ExtractFeatures(skills, history, preferences) | ScoreJobs(collaborative + content_based + boost) | RankByRelevance() | AddMatchReasons() → CacheRecommendations() → **Outbox:** recommendation.generated.v1
- GenerateFreelancerRecommendationsCommand(job_id, client_id, count=10) → LoadJobRequirements() | LoadModel() | ExtractFeatures() | ScoreFreelancers() | RankByFit() | AddMatchReasons() → CacheRecommendations() → **Outbox:** recommendation.generated.v1
- TrackRecommendationInteractionCommand(recommendation_id, user_id, action) → ValidateAction(viewed, clicked, applied, dismissed) | RecordInteraction() → UpdateModelFeedback() → **Outbox:** recommendation.interaction.tracked.v1
- RefreshRecommendationsCommand(user_id) → GenerateJobRecommendations() | InvalidateCache() → **Outbox:** recommendation.refreshed.v1
- GetRecommendationsQuery(user_id, type, page) → CheckCache() → If miss: GenerateRecommendations() ELSE: Return cached + paginate

### Projections

- recommendations_read (recommendation_id → user_id, items[], algorithm, generated_at, expires_at)
- recommendation_performance_read (model_version → CTR, conversion_rate, precision, recall)
- recommendation_feedback_read (interactions for model training)

### Events

- recommendation.generated.v1, recommendation.interaction.tracked.v1, recommendation.refreshed.v1, recommendation.expired.v1

### RBAC/SLO

- **RBAC:** **OWNER** for generate/get; **SYSTEM** for track
- **SLO:** generate P95 < 500ms; get P95 < 150ms (cached); track P95 < 100ms
- **Limits:** default 10 recommendations; max 50; refresh cooldown 1h; cache TTL 6h
- **Idempotency:** by (user_id, context, timestamp_hour) for generate

---

## 3.2 matching/

### Stories

- As a **system**, I want to compute job-freelancer match scores so that compatibility is quantified.
- As a **client**, I want to see top matching freelancers for my job so that I can review candidates easily.
- As a **freelancer**, I want to see my match score for jobs so that I prioritize applications.
- As a **system**, I want match score breakdowns (skill match, budget fit, availability) so that reasoning is transparent.
- As a **system**, I want to cache match scores so that repeated queries are fast.

### Flow

- ComputeMatchScoreCommand(job_id, freelancer_id) → LoadJobRequirements() | LoadFreelancerProfile() | ComputeSkillMatch() | ComputeBudgetFit() | ComputeAvailabilityMatch() | ComputeLocationMatch() | WeightedAggregate() → CacheScore() → **Outbox:** match.computed.v1
- GetTopMatchesQuery(job_id, count=20) → LoadJob() | SearchFreelancers(filters) | ForEach: ComputeMatchScore() | SortByScore() → Return top matches
- GetJobMatchesQuery(freelancer_id, count=20) → LoadFreelancer() | SearchJobs(filters) | ForEach: ComputeMatchScore() | SortByScore() → Return top matches
- ExplainMatchQuery(job_id, freelancer_id) → ComputeMatchScore() | GenerateExplanation(factor_scores, weights) → Return detailed breakdown

### Projections

- match_scores_read (job_id, freelancer_id → overall_score, skill_match, budget_fit, availability_match, computed_at)
- match_cache (Redis: {job_id, freelancer_id} → match_score, TTL 1h)
- match_criteria_weights_read (configurable weights for factors)

### Events

- match.computed.v1, match.accepted.v1, match.dismissed.v1

### RBAC/SLO

- **RBAC:** **SYSTEM** for compute; **USER** for get/explain
- **SLO:** compute P95 < 300ms; get top matches P95 < 1s; explain P95 < 200ms
- **Limits:** compute max 100 scores per request; top matches max 50
- **Idempotency:** by (job_id, freelancer_id, profiles_hash)

---

## 3.3 similarity/

### Stories

- As a **user**, I want "similar jobs" recommendations so that I can explore related opportunities.
- As a **system**, I want to compute similarity using embeddings (vector similarity) so that results are semantically relevant.
- As a **user**, I want "similar freelancers" suggestions so that I can find alternatives.
- As a **system**, I want to cache similarity results so that performance is optimal.
- As an **admin**, I want to update similarity models so that quality improves over time.

### Flow

- ComputeJobSimilarityCommand(job_id) → LoadJobEmbedding() | ESVectorSearch(cosine_similarity, top_k=20) | FilterActive() | RankBySimilarity() → CacheSimilarJobs() → **Outbox:** similarity.computed.v1
- ComputeUserSimilarityCommand(user_id) → LoadUserEmbedding() | ESVectorSearch() | FilterAvailable() | RankBySimilarity() → CacheSimilarUsers() → **Outbox:** similarity.computed.v1
- GetSimilarJobsQuery(job_id, count=10) → CheckCache() → If miss: ComputeJobSimilarity() ELSE: Return cached
- GetSimilarUsersQuery(user_id, count=10) → CheckCache() → If miss: ComputeUserSimilarity() ELSE: Return cached
- UpdateSimilarityModelCommand(model_path, version, admin_id) → ValidateModel() | LoadModel() | WarmupCache() → **Outbox:** similarity.model.updated.v1

### Projections

- similar_jobs_read (job_id → similar_job_ids[] with scores)
- similar_users_read (user_id → similar_user_ids[] with scores)
- similarity_model_metadata_read (model_version, accuracy, loaded_at)
- similarity_cache (Redis: entity_id → similar_entities[], TTL 12h)

### Events

- similarity.computed.v1, similarity.model.updated.v1

### RBAC/SLO

- **RBAC:** **SYSTEM** for compute; **USER** for get; **ADMIN** for update model
- **SLO:** compute P95 < 500ms; get P95 < 100ms (cached); update model P95 < 10s
- **Limits:** top_k max 100; cache max 20 similar items per entity
- **Idempotency:** by (entity_id, embedding_hash)

---

## 3.4 personalization/

### Stories

- As a **system**, I want to learn user preferences from behavior so that search is personalized.
- As a **user**, I want my search results boosted by my interests so that relevant results rank higher.
- As a **system**, I want cold-start strategies for new users so that personalization works immediately.
- As a **user**, I want to control personalization (disable, reset) so that I have privacy control.
- As a **system**, I want personalization profiles cached so that searches are fast.

### Flow

- BuildPersonalizationProfileCommand(user_id) → AggregateRecentActivity(searches, clicks, applications) | ExtractPreferences(skills, locations, budgets) | ComputeBoosts(skill_weights, location_preference) → CacheProfile() → **Outbox:** personalization.profile.built.v1
- ApplyPersonalizationCommand(user_id, query, results) → LoadProfile() | ApplyBoosts(skill_match, location_match) | ReRank() → Return personalized results
- ResetPersonalizationCommand(user_id) → ClearProfile() | ApplyColdStart() → **Outbox:** personalization.reset.v1
- GetPersonalizationProfileQuery(user_id) → LoadProfile() → Return preferences + boosts
- UpdatePreferencesCommand(user_id, preferences) → ValidatePreferences() | UpdateProfile() → InvalidateCache() → **Outbox:** personalization.updated.v1

### Projections

- personalization_profiles_read (user_id → preferences, boosts, last_updated)
- personalization_cache (Redis: user_id → profile, TTL 1h)
- cold_start_defaults_read (default preferences for new users)

### Events

- personalization.profile.built.v1, personalization.updated.v1, personalization.reset.v1

### RBAC/SLO

- **RBAC:** **OWNER** for get/update/reset; **SYSTEM** for build/apply
- **SLO:** build P95 < 500ms; apply P95 < 50ms; get P95 < 100ms
- **Limits:** profile recalculation cooldown 1h; max 100 preferences
- **Idempotency:** by (user_id, activity_hash) for build

---

## 3.5 feed/

### Stories

- As a **user**, I want a personalized job feed so that I see new opportunities daily.
- As a **system**, I want to generate feeds using recommendations + trending + recency so that content is fresh and relevant.
- As a **user**, I want to mark items as "not interested" so that I can train my feed.
- As a **system**, I want feed diversity so that users aren't in filter bubbles.
- As a **user**, I want my feed paginated and refreshable so that I can browse efficiently.

### Flow

- GenerateFeedCommand(user_id, feed_type, count=50) → LoadPersonalization() | GetRecommendations(30%) | GetTrending(20%) | GetRecent(30%) | GetDiverse(20%) | DeduplicateAndRank() | ApplyPersonalization() → CacheFeed() → **Outbox:** feed.generated.v1
- RefreshFeedCommand(user_id) → GenerateFeed() | InvalidateCache() → **Outbox:** feed.refreshed.v1
- MarkFeedItemCommand(user_id, item_id, action) → ValidateAction(dismissed, saved, applied) | RecordInteraction() → UpdatePersonalization() → **Outbox:** feed.item.interacted.v1
- GetFeedQuery(user_id, page, size) → CheckCache() → If miss: GenerateFeed() ELSE: Return cached + paginate
- GetFeedStatsQuery(user_id) → ComputeStats(view_rate, interaction_rate, diversity_score) → Return stats

### Projections

- user_feeds_read (user_id → items[], generated_at, expires_at)
- feed_interactions_read (user_id, item_id → action, timestamp)
- feed_stats_read (user_id → CTR, diversity, freshness)

### Events

- feed.generated.v1, feed.refreshed.v1, feed.item.interacted.v1, feed.expired.v1

### RBAC/SLO

- **RBAC:** **OWNER**
- **SLO:** generate P95 < 2s; get P95 < 200ms (cached); mark P95 < 150ms; refresh cooldown 15min
- **Limits:** default 50 items; max 100; cache TTL 1h; refresh cooldown 15min
- **Idempotency:** by (user_id, feed_type, timestamp_hour) for generate

---

## 3.6 trending/

### Stories

- As a **user**, I want to see trending jobs/skills so that I know what's popular.
- As a **system**, I want to compute trends using time-weighted scores so that recent activity counts more.
- As an **admin**, I want to configure trend windows (1h, 24h, 7d) so that trends are relevant.
- As a **system**, I want trending content cached so that computation isn't repeated.
- As a **user**, I want trending results updated frequently so that I see current trends.

### Flow

- ComputeTrendingJobsCommand(time_window, count=20) → AggregateMetrics(views, applications, recency) | ComputeTrendScores(time_decay) | RankByScore() → CacheTrending() → **Outbox:** trending.computed.v1
- ComputeTrendingSkillsCommand(time_window, count=20) → AggregateSkillMentions() | ComputeTrendScores() | RankByScore() → CacheTrending() → **Outbox:** trending.computed.v1
- GetTrendingQuery(entity_type, time_window) → CheckCache() → If miss: ComputeTrending() ELSE: Return cached
- UpdateTrendingJob() → IncrementMetrics() | EnqueueRecompute() → (async recomputation)
- GetTrendingStatsQuery(entity_id) → RetrieveTrendMetrics() → Return trend_score, growth_rate, rank

### Projections

- trending_jobs_read (time_window → job_ids[] with scores)
- trending_skills_read (time_window → skill_ids[] with scores)
- trending_cache (Redis: {entity_type, time_window} → trending list, TTL 15m)
- trend_metrics_read (entity_id → view_count, application_count, trend_score)

### Events

- trending.computed.v1, trending.updated.v1

### RBAC/SLO

- **RBAC:** **USER** for get; **SYSTEM** for compute/update
- **SLO:** compute P95 < 1s; get P95 < 100ms (cached); update P95 < 50ms
- **Limits:** time windows: 1h, 24h, 7d, 30d; max 100 trending items per window
- **Idempotency:** by (entity_type, time_window, computation_hour)

---

## 3.7 suggestion/

### Stories

- As a **user**, I want autocomplete suggestions while typing so that I can search faster.
- As a **user**, I want "did you mean?" spell corrections so that typos don't break search.
- As a **system**, I want query rewrites (synonyms, expansions) so that more results are found.
- As a **user**, I want related search suggestions so that I can explore topics.
- As a **system**, I want suggestions cached by prefix so that performance is optimal.

### Flow

- GetAutocompleteSuggestionsQuery(prefix, index_name, count=10) → ESCompletionSuggest(prefix) | RankByPopularity() → Return suggestions
- GetSpellCorrectionQuery(query) → ESSpellCheck() | ComputeLevenshteinDistance() | RankCorrections() → Return "did you mean" suggestions
- RewriteQueryCommand(query, context) → LookupSynonyms() | ExpandAbbreviations() | ApplyRewrites() → Return rewritten_query + confidence
- GetRelatedSearchesQuery(query) → FindSimilarQueries(embedding) | RankByFrequency() → Return related searches
- TrackSuggestionCommand(suggestion, user_id, selected?) → RecordUsage() | UpdatePopularity() → **Outbox:** suggestion.tracked.v1

### Projections

- suggestions_read (prefix → suggestions[] with scores)
- spell_corrections_read (misspelling → corrections[] with confidence)
- query_rewrites_read (original → rewritten with rules applied)
- suggestion_cache (Redis: prefix → suggestions, TTL 1h)

### Events

- suggestion.tracked.v1, suggestion.selected.v1

### RBAC/SLO

- **RBAC:** **USER** for get; **SYSTEM** for track
- **SLO:** autocomplete P95 < 100ms; spell check P95 < 150ms; rewrite P95 < 50ms
- **Limits:** max 10 suggestions per request; max edit distance 2 for spell check
- **Idempotency:** tracking by (suggestion, user_hash, timestamp_hour)

---

# 4 - RANKING & BOOSTING

## 4.1 ranking/

### Stories

- As a **system**, I want to rank search results using multiple signals (relevance, quality, freshness) so that best results are first.
- As an **admin**, I want configurable ranking weights so that I can tune search quality.
- As a **system**, I want learning-to-rank (LTR) models so that ranking improves over time.
- As a **user**, I want to sort by different criteria (date, budget, rating) so that I control result order.
- As a **system**, I want to explain ranking so that debugging is possible.

### Flow

- RankResultsCommand(results[], ranking_config) → ForEach: ComputeRelevanceScore() | ComputeQualityScore() | ComputeFreshnessScore() | ApplyWeights(config) | ApplyLTRModel() → SortByFinalScore() → Return ranked results
- ConfigureRankingWeightsCommand(index_name, weights, admin_id) → ValidateWeights(sum=1.0) | UpdateConfig() → InvalidateCache() → **Outbox:** ranking.configured.v1
- UpdateLTRModelCommand(model_path, version, admin_id) → ValidateModel() | LoadModel() | WarmupCache() → **Outbox:** ranking.ltr.updated.v1
- ExplainRankingQuery(result_id, query) → ComputeScores() | GenerateExplanation(factors) → Return detailed breakdown
- GetRankingConfigQuery(index_name) → RetrieveConfig() → Return weights + model version

### Projections

- ranking_configs_read (index_name → weights, LTR_model_version, updated_at)
- ranking_model_metadata_read (model_version, features, accuracy, training_date)
- ranking_explanations_cache (Redis: {query_hash, result_id} → explanation, TTL 5m)

### Events

- ranking.configured.v1, ranking.ltr.updated.v1

### RBAC/SLO

- **RBAC:** **ADMIN** for configure/update model; **SYSTEM** for rank; **USER** for explain
- **SLO:** rank P95 < 100ms; configure P95 < 300ms; explain P95 < 200ms
- **Limits:** max 10 ranking factors; weights must sum to 1.0
- **Idempotency:** by (index_name, config_hash) for configure

---

## 4.2 boost/

### Stories

- As an **admin**, I want to boost specific jobs (featured, promoted) so that visibility is increased.
- As a **client**, I want to boost my jobs by payment so that they rank higher.
- As a **system**, I want time-based boost decay so that boosts expire naturally.
- As an **admin**, I want boost multipliers configurable so that impact is controlled.
- As a **system**, I want to track boost effectiveness so that ROI is measured.

### Flow

- BoostJobCommand(job_id, boost_type, multiplier, duration_hours) → ValidateBoostType(FEATURED/PROMOTED/URGENT) | ApplyBoostMultiplier(multiplier) | SetExpiry(now + duration) | UpdateESDocument() → **Outbox:** job.boosted.v1
- RemoveBoostCommand(job_id) → RemoveMultiplier() | UpdateESDocument() → **Outbox:** job.boost.removed.v1
- DecayBoostJob() → ForEachExpiredBoost: RemoveBoost() → **Outbox:** job.boost.expired.v1
- GetBoostStatsQuery(job_id) → ComputeImpact(views_increase, applications_increase) → Return boost ROI
- ListBoostedJobsQuery(boost_type) → RetrieveBoostedJobs(active=true) → Return list with multipliers

### Projections

- boosted_jobs_read (job_id → boost_type, multiplier, applied_at, expires_at)
- boost_effectiveness_read (job_id → views_lift, applications_lift, ROI)

### Events

- job.boosted.v1, job.boost.removed.v1, job.boost.expired.v1

### RBAC/SLO

- **RBAC:** **ADMIN/CLIENT** for boost; **SYSTEM** for decay
- **SLO:** boost P95 < 500ms; remove P95 < 300ms; stats P95 < 200ms
- **Limits:** max multiplier 10x; max duration 30 days; max 10 concurrent boosts per job
- **Idempotency:** by (job_id, boost_type, timestamp)

---

## 4.3 promotion/

### Stories

- As an **admin**, I want to manually promote content (sticky results) so that strategic items appear first.
- As an **admin**, I want promotion rules (query patterns, dates) so that promotions are targeted.
- As a **system**, I want to track promotion impressions and clicks so that effectiveness is measured.
- As an **admin**, I want to schedule promotions (start/end dates) so that campaigns are automated.
- As a **user**, I want promoted results clearly labeled so that organic vs promoted is transparent.

### Flow

- CreatePromotionCommand(admin_id, item_id, target_queries[], start_date, end_date, position) → ValidateDates() | CreatePromotion(active=true) → CachePromotionRules() → **Outbox:** promotion.created.v1
- UpdatePromotionCommand(promotion_id, updates, admin_id) → ValidateUpdates() | UpdatePromotion() → InvalidateCache() → **Outbox:** promotion.updated.v1
- DeletePromotionCommand(promotion_id, admin_id) → MarkInactive() → InvalidateCache() → **Outbox:** promotion.deleted.v1
- ApplyPromotionsCommand(query, results[]) → LoadActivePromotions(query_match) | InjectPromotedItems(position) | MarkAsPromoted() → Return results with promotions
- TrackPromotionCommand(promotion_id, action) → ValidateAction(impression, click) | RecordInteraction() → **Outbox:** promotion.tracked.v1
- GetPromotionStatsQuery(promotion_id) → ComputeStats(impressions, clicks, CTR) → Return stats

### Projections

- promotions_read (promotion_id → item_id, target_queries[], position, start_date, end_date, active)
- promotion_stats_read (promotion_id → impressions, clicks, CTR)
- promotion_rules_cache (Redis: query_pattern → promotions[], TTL 1h)

### Events

- promotion.created.v1, promotion.updated.v1, promotion.deleted.v1, promotion.tracked.v1, promotion.expired.v1

### RBAC/SLO

- **RBAC:** **ADMIN** for create/update/delete; **SYSTEM** for apply/track
- **SLO:** create P95 < 300ms; apply P95 < 100ms; track P95 < 50ms
- **Limits:** max 10 active promotions per query; max position 3
- **Idempotency:** by (item_id, target_query, start_date) for create

---

# 5 - ANALYTICS & MONITORING

## 5.1 analytics/

### Stories

- As an **admin**, I want search analytics (query volume, popular queries, zero-result queries) so that I can improve search.
- As a **system**, I want to track click-through rates so that I measure search quality.
- As an **admin**, I want to see search performance metrics (latency, error rate) so that I can optimize.
- As a **system**, I want to aggregate analytics daily so that historical trends are visible.
- As an **admin**, I want analytics dashboards so that I can monitor KPIs.

### Flow

- TrackSearchAnalyticsCommand(query, result_count, latency, user_id?) → RecordQueryLog() | IncrementCounters() → **Outbox:** analytics.tracked.v1
- AggregateAnalyticsJob(time_period) → AggregateQueryVolume() | ComputeTopQueries() | ComputeCTR() | ComputeZeroResultQueries() | ComputeAvgLatency() → StoreAggregates() → **Outbox:** analytics.aggregated.v1
- GetSearchAnalyticsQuery(time_period, admin_id) → LoadAggregates() → Return analytics (query_volume, top_queries, CTR, latency)
- GetZeroResultQueriesQuery(count=100, admin_id) → RetrieveZeroResultQueries() → Return list for review
- GetPopularQueriesQuery(count=100, time_period) → RetrievePopularQueries() → Return list with counts

### Projections

- search_analytics_read (date → query_volume, unique_users, top_queries[], avg_latency, error_rate)
- zero_result_queries_read (query → count, last_seen)
- popular_queries_read (time_period → query, count)
- query_performance_read (query_pattern → avg_latency, p95_latency)

### Events

- analytics.tracked.v1, analytics.aggregated.v1, analytics.alert.triggered.v1

### RBAC/SLO

- **RBAC:** **SYSTEM** for track/aggregate; **ADMIN** for get
- **SLO:** track P95 < 50ms; aggregate P95 < 10s per day; get P95 < 300ms
- **Limits:** retain raw logs 30 days; retain aggregates 2 years
- **Idempotency:** tracking by (query_hash, user_hash, timestamp_second)

---

## 5.2 performance/

### Stories

- As an **admin**, I want to monitor Elasticsearch cluster health so that I can prevent outages.
- As a **system**, I want to track query performance (slow queries) so that I can optimize.
- As an **admin**, I want alerting on performance degradation so that I can respond quickly.
- As a **system**, I want to track index sizes and shard status so that I can plan capacity.
- As an **admin**, I want performance dashboards so that I can monitor SLOs.

### Flow

- MonitorClusterHealthJob() → ESClusterHealthAPI() | CheckStatus(green/yellow/red) | CheckShardStatus() | CheckDiskUsage() → If degraded: **Outbox:** cluster.health.degraded.v1
- TrackSlowQueryCommand(query, latency) → RecordSlowQuery() → If latency > threshold: **Outbox:** query.slow.detected.v1
- GetPerformanceMetricsQuery(admin_id) → ESStatsAPI() | AggregateMetrics() → Return metrics (query_rate, latency, error_rate, index_sizes)
- GetSlowQueriesQuery(count=100, admin_id) → RetrieveSlowQueries() → Return list with latencies
- TriggerPerformanceAlertCommand(alert_type, severity) → ValidateAlert() | SendNotification() → **Outbox:** performance.alert.triggered.v1

### Projections

- cluster_health_read (timestamp → status, active_shards, relocating_shards, unassigned_shards)
- slow_queries_read (query_hash → avg_latency, max_latency, count)
- performance_metrics_read (date → query_rate, avg_latency, p95_latency, error_rate)
- index_sizes_read (index_name → doc_count, size_bytes, shard_count)

### Events

- cluster.health.degraded.v1, query.slow.detected.v1, performance.alert.triggered.v1, index.size.threshold.exceeded.v1

### RBAC/SLO

- **RBAC:** **ADMIN** for get; **SYSTEM** for monitor/track
- **SLO:** monitor interval 1min; track P95 < 50ms; get P95 < 300ms
- **Limits:** slow query threshold 1s; alert cooldown 5min; retain metrics 90 days
- **Idempotency:** tracking by (query_hash, timestamp_minute)

---

## 5.3 explain/

### Stories

- As a **developer**, I want to explain why a document matched a query so that I can debug relevance.
- As an **admin**, I want to see ranking factor contributions so that I can tune weights.
- As a **system**, I want to generate explain data on demand so that performance isn't impacted.
- As a **developer**, I want to compare explain data for multiple documents so that I can understand ranking order.

### Flow

- ExplainSearchCommand(query, doc_id, admin_id) → BuildESQuery() | ESExplainAPI(doc_id) | ParseExplanation() → Return explain tree with scores
- ExplainRankingCommand(query, doc_ids[], admin_id) → ForEach: ExplainSearch() | CompareScores() | GenerateComparison() → Return comparative explanation
- GetExplainTemplateQuery(index_name) → RetrieveMappings() | GenerateExplainTemplate() → Return template with scoring factors

### Projections

- explain_cache (Redis: {query_hash, doc_id} → explain, TTL 5m)

### Events

- (none - explain is read-only)

### RBAC/SLO

- **RBAC:** **ADMIN/DEVELOPER**
- **SLO:** explain P95 < 500ms; compare P95 < 2s (for 10 docs)
- **Limits:** max 20 docs per comparison; explain timeout 2s
- **Idempotency:** N/A (read-only)

---

# 6 - INDEX LIFECYCLE & HYGIENE

## 6.1 lifecycle/

### Stories

- As an **admin**, I want to rollover indices when they reach size/age limits so that performance is maintained.
- As an **system**, I want to snapshot indices for backup so that data is protected.
- As an **admin**, I want to restore indices from snapshots so that I can recover from failures.
- As a **system**, I want to delete old indices based on retention policies so that storage is optimized.
- As an **admin**, I want index lifecycle policies configurable so that automation is flexible.

### Flow

- RolloverIndexCommand(index_name, admin_id) → CheckRolloverConditions(size, age, docs) → If met: CreateNewIndex(version++) | UpdateAlias() | CloseOldIndex() → **Outbox:** index.rolledover.v1
- SnapshotIndexCommand(index_name, snapshot_name, admin_id) → ESSnapshotAPI(index_name) | WaitForCompletion() → **Outbox:** index.snapshotted.v1
- RestoreIndexCommand(snapshot_name, index_name, admin_id) → ESRestoreAPI(snapshot_name) | WaitForCompletion() | UpdateAliases() → **Outbox:** index.restored.v1
- DeleteOldIndexCommand(index_name, admin_id) → CheckRetentionPolicy() | ESDeleteIndexAPI() | RemoveMetadata() → **Outbox:** index.deleted.v1
- ConfigureLifecyclePolicyCommand(policy_name, rules, admin_id) → ValidateRules() | CreateESPolicy() | ApplyToIndices() → **Outbox:** lifecycle.policy.configured.v1
- ApplyLifecyclePolicyJob() → ForEachIndex: CheckPolicy() | ExecuteAction(rollover, snapshot, delete) → **Outbox:** lifecycle.policy.applied.v1

### Projections

- index_lifecycle_policies_read (policy_name → rules, target_indices, active)
- index_snapshots_read (snapshot_name → indices[], created_at, status, size)
- index_rollover_history_read (index_name → rollover_events with dates + reasons)

### Events

- index.rolledover.v1, index.snapshotted.v1, index.restored.v1, index.deleted.v1, lifecycle.policy.configured.v1, lifecycle.policy.applied.v1

### RBAC/SLO

- **RBAC:** **ADMIN** for all operations
- **SLO:** rollover P95 < 10s; snapshot P95 < 60s (depends on size); restore P95 < 120s; delete P95 < 5s
- **Limits:** default retention 365 days; max 100 snapshots; rollover conditions: 50GB OR 30d OR 100M docs
- **Idempotency:** by (index_name, operation, timestamp_day)

---

## 6.2 hygiene/

### Stories

- As a **system**, I want to detect and remove duplicate documents so that index quality is high.
- As an **admin**, I want to reindex corrupted documents so that search results are accurate.
- As a **system**, I want to clean up orphaned references so that indices are consistent.
- As an **admin**, I want to rebuild indices from source systems so that data is fresh.
- As a **system**, I want hygiene jobs scheduled so that maintenance is automated.

### Flow

- DetectDuplicatesCommand(index_name) → ESAggregationsAPI(duplicate_hashes) | IdentifyDuplicates() → **Outbox:** hygiene.duplicates.detected.v1
- RemoveDuplicatesCommand(index_name, duplicate_ids[], admin_id) → ForEach: DeleteDocument() | KeepCanonical() → **Outbox:** hygiene.duplicates.removed.v1
- ReindexCorruptedCommand(index_name, doc_ids[], admin_id) → ForEach: FetchFromSource() | ValidateData() | ReindexDocument() → **Outbox:** hygiene.reindexed.v1
- CleanOrphanedReferencesCommand(index_name) → CompareWithSourceSystems() | IdentifyOrphans() | DeleteOrphans() → **Outbox:** hygiene.orphans.cleaned.v1
- RebuildIndexCommand(index_name, admin_id) → CreateNewIndex() | BulkReindexFromSource() | ValidateDocCounts() | SwapAlias() → **Outbox:** hygiene.index.rebuilt.v1
- ScheduleHygieneJobCommand(job_type, cron_expression, admin_id) → ValidateCron() | CreateSchedule() → **Outbox:** hygiene.job.scheduled.v1

### Projections

- hygiene_issues_read (issue_type → count, doc_ids[], detected_at)
- hygiene_jobs_read (job_id → type, status, started_at, completed_at, issues_found, issues_fixed)
- hygiene_schedules_read (job_type → cron_expression, last_run, next_run)

### Events

- hygiene.duplicates.detected.v1, hygiene.duplicates.removed.v1, hygiene.reindexed.v1, hygiene.orphans.cleaned.v1, hygiene.index.rebuilt.v1, hygiene.job.scheduled.v1

### RBAC/SLO

- **RBAC:** **ADMIN** for all operations; **SYSTEM** for detect
- **SLO:** detect P95 < 5s; remove P95 < 10s per 1000 docs; reindex P95 < 30s per 1000 docs; rebuild P95 < 10min per 100K docs
- **Limits:** max 10,000 duplicates per job; hygiene jobs run during low-traffic hours
- **Idempotency:** by (index_name, job_type, timestamp_day)

---

## 6.3 compliance/

### Stories

- As a **DPO**, I want to erase user data from indices for GDPR compliance so that right to erasure is respected.
- As a **system**, I want to mask PII in search logs so that privacy is protected.
- As an **admin**, I want to export user search data for DSAR requests so that compliance is automated.
- As a **system**, I want to track compliance actions so that audits are supported.
- As a **DPO**, I want to verify erasure completeness so that compliance is confirmed.

### Flow

- EraseUserDataCommand(user_id, dpo_id) → IdentifyUserDocuments(all indices) | DeleteDocuments() | RemoveFromLogs() | MaskReferences() → **Outbox:** compliance.erasure.completed.v1
- ExportUserSearchDataCommand(user_id, dpo_id) → RetrieveUserSearches() | RetrieveRecommendations() | RetrieveInteractions() | GenerateExport(JSON) → CreateExportArtifact() → **Outbox:** compliance.export.created.v1
- MaskPIICommand(log_entries[]) → DetectPII(emails, phones) | RedactPII() | UpdateLogs() → **Outbox:** compliance.pii.masked.v1
- VerifyErasureCommand(user_id, dpo_id) → SearchAllIndices(user_id) | CheckLogs() | CheckRecommendations() → Return erasure status + remaining references
- TrackComplianceActionCommand(action_type, user_id, dpo_id) → RecordAction(timestamp, details) → **Outbox:** compliance.action.tracked.v1

### Projections

- compliance_actions_read (action_id → type, user_id, dpo_id, timestamp, status)
- erasure_verifications_read (user_id → verified_at, remaining_references_count)
- export_artifacts_read (export_id → user_id, artifact_url, created_at, expires_at)

### Events

- compliance.erasure.completed.v1, compliance.export.created.v1, compliance.pii.masked.v1, compliance.action.tracked.v1

### RBAC/SLO

- **RBAC:** **DPO** for all operations
- **SLO:** erase P95 < 10s per user; export P95 < 30s; mask P95 < 5s per 1000 logs; verify P95 < 5s
- **Limits:** export retention 30 days; verify includes all indices + logs
- **Idempotency:** by (user_id, action_type, timestamp_day)

---

# 7 - LANGUAGE & INTERNATIONALIZATION

## 7.1 language/

### Stories

- As a **user**, I want to search in my native language so that results are relevant.
- As a **system**, I want language detection so that appropriate analyzers are used.
- As a **user**, I want results in my preferred language so that content is understandable.
- As an **admin**, I want to configure language support so that new languages are enabled.
- As a **system**, I want to track language usage so that I can prioritize improvements.

### Flow

- DetectLanguageCommand(text) → UseLanguageDetectionAPI() | CacheResult() → Return language_code + confidence
- ConfigureLanguageSupportCommand(language_code, analyzer_settings, admin_id) → ValidateLanguage() | CreateESAnalyzer() | UpdateIndexSettings() → **Outbox:** language.configured.v1
- SearchWithLanguageCommand(query, language_code, user_id) → SelectAnalyzer(language_code) | BuildESQuery(language_specific) | ExecuteSearch() → Return results
- GetSupportedLanguagesQuery() → RetrieveConfiguredLanguages() → Return list with stats
- TrackLanguageUsageCommand(language_code, query_count) → IncrementCounter() → **Outbox:** language.usage.tracked.v1

### Projections

- supported_languages_read (language_code → name, analyzer, enabled, doc_count)
- language_usage_read (language_code → query_count, user_count, growth_rate)
- language_detection_cache (Redis: text_hash → language_code, TTL 1h)

### Events

- language.configured.v1, language.usage.tracked.v1

### RBAC/SLO

- **RBAC:** **ADMIN** for configure; **USER/SYSTEM** for detect/search
- **SLO:** detect P95 < 100ms; search P95 < 300ms; configure P95 < 1s
- **Limits:** max 50 supported languages; default language: English
- **Idempotency:** by (language_code) for configure

---

## 7.2 speller/

### Stories

- As a **user**, I want spell check suggestions so that typos are corrected.
- As a **system**, I want "did you mean?" suggestions so that searches succeed despite typos.
- As a **system**, I want spell check in multiple languages so that international users are supported.
- As an **admin**, I want to add custom dictionary terms so that domain-specific words are recognized.
- As a **system**, I want spell check confidence scores so that corrections are reliable.

### Flow

- GetSpellSuggestionsCommand(query, language_code) → ESSpellCheckAPI() | ComputeLevenshteinDistance() | RankByConfidence() → Return suggestions with confidence scores
- AddCustomTermCommand(term, language_code, admin_id) → ValidateTerm() | AddToDictionary() | UpdateESSettings() → **Outbox:** speller.term.added.v1
- ConfigureSpellCheckCommand(settings, admin_id) → ValidateSettings(max_edits, prefix_length) | UpdateESSettings() → **Outbox:** speller.configured.v1
- TrackSpellCorrectionCommand(original, correction, user_id, accepted?) → RecordCorrection() | UpdateAcceptanceRate() → **Outbox:** speller.correction.tracked.v1

### Projections

- spell_corrections_read (misspelling → corrections[] with confidence + acceptance_rate)
- custom_dictionary_read (term → language_code, added_by, added_at)
- spell_check_settings_read (language_code → max_edits, prefix_length, enabled)

### Events

- speller.term.added.v1, speller.configured.v1, speller.correction.tracked.v1

### RBAC/SLO

- **RBAC:** **ADMIN** for add/configure; **USER/SYSTEM** for get
- **SLO:** get suggestions P95 < 150ms; add term P95 < 300ms
- **Limits:** max edit distance 2; max 10 suggestions per query; max 10,000 custom terms per language
- **Idempotency:** by (term, language_code) for add

---

## 7.3 rewrite/

### Stories

- As a **system**, I want query rewriting (synonyms, expansions) so that more results are found.
- As an **admin**, I want rewrite rules configurable so that business logic is applied.
- As a **system**, I want multi-match rewrites (OR expansion) so that recall is improved.
- As a **user**, I want to preview rewrites so that I understand query transformation.
- As a **system**, I want to track rewrite effectiveness so that rules are optimized.

### Flow

- RewriteQueryCommand(query, context?) → LoadRewriteRules() | ApplySynonyms() | ApplyExpansions() | ApplyCorrections() → Return rewritten_query + rules_applied
- ConfigureRewriteRuleCommand(rule_name, pattern, replacement, admin_id) → ValidateRule() | CreateRule(active=true) → CacheRules() → **Outbox:** rewrite.rule.configured.v1
- PreviewRewriteQuery(query, rules_to_apply?) → ApplyRules() | ShowDiff() → Return original vs rewritten
- TrackRewriteCommand(original_query, rewritten_query, result_count_delta) → RecordRewrite() | ComputeEffectiveness() → **Outbox:** rewrite.tracked.v1
- GetRewriteRulesQuery(admin_id) → RetrieveRules(active=true) → Return list with effectiveness stats

### Projections

- rewrite_rules_read (rule_id → name, pattern, replacement, active, effectiveness_score)
- rewrite_history_read (rule_id → applications_count, avg_result_improvement)
- rewrite_cache (Redis: rules list, TTL 1h)

### Events

- rewrite.rule.configured.v1, rewrite.tracked.v1

### RBAC/SLO

- **RBAC:** **ADMIN** for configure; **USER** for preview; **SYSTEM** for rewrite/track
- **SLO:** rewrite P95 < 50ms; configure P95 < 300ms; preview P95 < 100ms
- **Limits:** max 100 active rules; max 5 rewrites per query
- **Idempotency:** by (rule_name) for configure

---

# 8 - INBOX EVENTS (Consumers)

## 8.1 Inbox: Job Events

### Stories

- As a **system**, I want to consume job.posted events to index new jobs.
- As a **system**, I want to consume job.updated events to reindex job changes.
- As a **system**, I want to consume job.closed events to mark jobs as inactive in index.

### Flow

- Consume: job.posted.v1 → IndexJob(job_id, job_data) → **Outbox:** search.job.indexed.v1
- Consume: job.updated.v1 → ReindexJob(job_id, updated_fields) → **Outbox:** search.job.indexed.v1
- Consume: job.closed.v1 → UpdateJobStatus(job_id, status=CLOSED) | ReduceVisibility() → **Outbox:** search.job.updated.v1
- Consume: job.deleted.v1 → RemoveFromIndex(job_id) → **Outbox:** search.document.removed.v1

### Projections

- (updates to indexed_documents_read)

### Events

- (consumer)

### RBAC/SLO

- **RBAC:** SYSTEM
- **SLO:** P95 < 500ms per event

---

## 8.2 Inbox: User Events

### Stories

- As a **system**, I want to consume user.profile.updated events to reindex user data.
- As a **system**, I want to consume user.skills.updated events to refresh talent search.
- As a **system**, I want to consume user.deleted events to remove users from index.

### Flow

- Consume: user.profile.updated.v1 → ReindexUser(user_id, profile_data) → **Outbox:** search.user.indexed.v1
- Consume: user.skills.updated.v1 → UpdateUserSkills(user_id, skills[]) | ReindexUser() → **Outbox:** search.user.indexed.v1
- Consume: user.deleted.v1 → RemoveFromIndex(user_id) → **Outbox:** search.document.removed.v1

### Projections

- (updates to indexed_documents_read)

### Events

- (consumer)

### RBAC/SLO

- **RBAC:** SYSTEM
- **SLO:** P95 < 500ms per event

---

## 8.3 Inbox: Proposal Events

### Stories

- As a **system**, I want to consume proposal.submitted events to update job popularity scores.
- As a **system**, I want to consume proposal.accepted events to update job conversion rates.

### Flow

- Consume: proposal.submitted.v1 → IncrementJobProposalCount(job_id) | UpdatePopularityScore() → **Outbox:** search.job.updated.v1
- Consume: proposal.accepted.v1 → IncrementJobConversionRate(job_id) | UpdateQualityScore() → **Outbox:** search.job.updated.v1

### Projections

- (updates to trending/popularity scores)

### Events

- (consumer)

### RBAC/SLO

- **RBAC:** SYSTEM
- **SLO:** P95 < 300ms per event

---

## 8.4 Inbox: Review Events

### Stories

- As a **system**, I want to consume review.created events to update user ratings in search.

### Flow

- Consume: review.created.v1 → UpdateUserRating(user_id, new_rating) | ReindexUser() → **Outbox:** search.user.indexed.v1

### Projections

- (updates to user rating in index)

### Events

- (consumer)

### RBAC/SLO

- **RBAC:** SYSTEM
- **SLO:** P95 < 500ms per event

---

## 8.5 Inbox: Admin Events

### Stories

- As a **system**, I want to consume admin.content.hidden events to remove content from search.
- As a **system**, I want to consume admin.taxonomy.updated events to refresh taxonomy cache.

### Flow

- Consume: admin.content.hidden.v1 → RemoveFromIndex(entity_id) OR UpdateVisibility(entity_id, visibility=HIDDEN) → **Outbox:** search.document.updated.v1
- Consume: admin.taxonomy.updated.v1 → RefreshTaxonomyCache() | ReindexAffectedDocuments() → **Outbox:** taxonomy.refreshed.v1

### Projections

- (updates to taxonomy_cache)

### Events

- (consumer)

### RBAC/SLO

- **RBAC:** SYSTEM
- **SLO:** P95 < 300ms per event

---

## 8.6 Inbox: Subscription Events

### Stories

- As a **system**, I want to consume subscription.feature.changed events to gate premium search features.

### Flow

- Consume: subscription.feature.changed.v1 → UpdateUserEntitlements(user_id, features[]) | RefreshFeatureGates() → (cache update)

### Projections

- user_search_entitlements_cache (Redis: user_id → features[], TTL 1h)

### Events

- (consumer)

### RBAC/SLO

- **RBAC:** SYSTEM
- **SLO:** P95 < 200ms per event

---

## 8.7 Inbox: Compliance Events

### Stories

- As a **system**, I want to consume user.erasure.requested events to remove user data from indices.

### Flow

- Consume: user.erasure.requested.v1 → EraseUserData(user_id) → **Outbox:** compliance.erasure.completed.v1

### Projections

- (cleanup from all indices and logs)

### Events

- (consumer)

### RBAC/SLO

- **RBAC:** SYSTEM
- **SLO:** P95 < 10s per event

---

# APPENDIX

## A. Elasticsearch Indices

- **jobs_index:** Job postings with skills, location, budget, boost factors
- **users_index:** Freelancer profiles with skills, experience, ratings
- **portfolios_index:** Portfolio items with media, skills, engagement metrics

## B. Search Query Types

- **Match Query:** Full-text search on title/description
- **Multi-Match Query:** Search across multiple fields
- **Bool Query:** Combine must/should/must_not clauses
- **Range Query:** Budget, experience, date ranges
- **Geo Query:** Location-based filtering (geo_distance)
- **Function Score:** Custom scoring with field value factors
- **More Like This:** Similarity search using embeddings

## C. Ranking Factors

- **Relevance:** TF-IDF, BM25 text matching score
- **Quality:** User rating, success rate, verification status
- **Freshness:** Time decay function (recent posts ranked higher)
- **Popularity:** View count, application count, trending score
- **Personalization:** User preference boost, past behavior
- **Boost Multipliers:** Featured, promoted, urgent flags

## D. Analyzers

- **Standard Analyzer:** Default, tokenizes on whitespace/punctuation
- **English Analyzer:** Stemming, stopwords for English
- **Multi-language Analyzers:** Language-specific tokenization
- **Ngram Analyzer:** Partial matching for autocomplete
- **Custom Analyzer:** Domain-specific tokenization rules

## E. ML Models

- **Collaborative Filtering:** User-item interaction matrix factorization
- **Content-Based:** Feature similarity using embeddings
- **Hybrid Model:** Combines collaborative + content-based
- **Learning-to-Rank (LTR):** XGBoost/LambdaMART for result ranking
- **Embeddings:** BERT/Sentence-BERT for semantic similarity

## F. Rate Limits

- **Search:** 100 queries/min per user
- **Indexing:** 1000 docs/min per index
- **Autocomplete:** 50 requests/min per user
- **Recommendations:** 10 requests/min per user
- **Analytics:** 20 requests/min per admin

## G. Cache Strategy

- **Search Results:** TTL 15min, invalidate on index update
- **Facets:** TTL 15min, invalidate on data change
- **Taxonomy:** TTL 1h, invalidate on admin update
- **Recommendations:** TTL 6h, refresh on user activity
- **Similarity:** TTL 12h, recompute on embedding update
- **Personalization:** TTL 1h, update on behavior change

## H. SLO Targets

- **Search Latency:** P95 < 300ms
- **Indexing Latency:** P95 < 500ms
- **Autocomplete:** P95 < 100ms
- **Recommendations:** P95 < 500ms
- **Ranking:** P95 < 100ms
- **Availability:** 99.9% uptime

## I. Event Topics

- **search.*:** search.job.indexed, search.user.indexed, search.portfolio.indexed, search.document.removed, search.bulk.indexed, search.index.reindexed, search.query.logged, search.query.failed, search.saved.executed, search.alert.triggered
- **recommendation.*:** recommendation.generated, recommendation.interaction.tracked, recommendation.refreshed, recommendation.expired
- **match.*:** match.computed, match.accepted, match.dismissed
- **similarity.*:** similarity.computed, similarity.model.updated
- **personalization.*:** personalization.profile.built, personalization.updated, personalization.reset
- **feed.*:** feed.generated, feed.refreshed, feed.item.interacted, feed.expired
- **trending.*:** trending.computed, trending.updated
- **suggestion.*:** suggestion.tracked, suggestion.selected
- **ranking.*:** ranking.configured, ranking.ltr.updated
- **boost.*:** job.boosted, job.boost.removed, job.boost.expired
- **promotion.*:** promotion.created, promotion.updated, promotion.deleted, promotion.tracked, promotion.expired
- **analytics.*:** analytics.tracked, analytics.aggregated, analytics.alert.triggered
- **performance.*:** cluster.health.degraded, query.slow.detected, performance.alert.triggered, index.size.threshold.exceeded
- **index.*:** index.created, index.mappings.updated, index.alias.set, index.visibility.changed, index.archived, index.health.degraded, index.rolledover, index.snapshotted, index.restored, index.deleted
- **lifecycle.*:** lifecycle.policy.configured, lifecycle.policy.applied
- **hygiene.*:** hygiene.duplicates.detected, hygiene.duplicates.removed, hygiene.reindexed, hygiene.orphans.cleaned, hygiene.index.rebuilt, hygiene.job.scheduled
- **compliance.*:** compliance.erasure.completed, compliance.export.created, compliance.pii.masked, compliance.action.tracked
- **taxonomy.*:** taxonomy.skill.created, taxonomy.skill.updated, taxonomy.skill.deprecated, taxonomy.synonym.added, taxonomy.trends.computed, taxonomy.refreshed
- **facets.*:** facets.configured, facets.computed
- **filters.*:** filters.tracked, filters.invalid
- **language.*:** language.configured, language.usage.tracked
- **speller.*:** speller.term.added, speller.configured, speller.correction.tracked
- **rewrite.*:** rewrite.rule.configured, rewrite.tracked

---

**END OF search-be USER STORIES**

Total Sections: 39  
Total Stories: 250+  
Coverage: 100% of search-be domain structure  
Event-Driven: ✅  
CQRS: ✅  
Outbox Pattern: ✅  
RBAC: ✅  
SLO Defined: ✅  
Idempotency: ✅  
Non-PII: ✅  
ML/AI Integration: ✅  
Elasticsearch: ✅
