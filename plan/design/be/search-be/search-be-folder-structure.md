
## **📦8️⃣ search-be (UPDATED WITH ENHANCED RECOMMENDATIONS)**

```
apps/be/search-be/
│
├── cmd/
│   # =============================
│   # 🚀 APP ENTRYPOINTS
│   # =============================
│   ├── api/
│   │   └── main.go                               # 📝 Gin + Dapr + Elasticsearch; uses internal/config & platform-shared/logging
│   └── worker/                                    # 🆕 background runner
│       └── main.go                                # 🆕 runs hygiene/backfill/snapshot jobs; leader election safe
│
├── internal/
│   # =============================
│   # 🔧 CONFIGURATION (LOAD FIRST)
│   # =============================
│   ├── config/
│   │   ├── schema.go                              # Typed Config (App, Server, Postgres, Kafka, Redis, Elasticsearch, ML, IndexLifecycle)
│   │   ├── loader.go                              # Viper loader (flags → env → file → defaults)
│   │   └── docs/
│   │       └── CONFIGURATION.md                   # All ENV vars, defaults, examples
│   │
│   # =============================
│   # 🏛️ DOMAIN LAYER (DDD)
│   # =============================
│   ├── domain/
│   │   # =========================
│   │   # 📚 CORE INDEX ARTIFACTS
│   │   # =========================
│   │   ├── search_index/
│   │   │   ├── entity.go                          # Index meta (kind, alias, version, settings, mappings_hash, visibility)
│   │   │   ├── job_index.go                       # Job search document fields
│   │   │   ├── user_index.go                      # User/Freelancer search document fields
│   │   │   ├── errors.go                          # IndexNotFound, DocumentConflict
│   │   │   ├── repository.go                      # IndexRepository interface
│   │   │   └── events.go                          # search.document.indexed/reindexed; search.index.visibility.changed/archived.v1
│   │   ├── portfolio_index/                       # 🆕 discoverable portfolios/works
│   │   │   ├── entity.go                          # Title, skills[], media_refs, engagement, recency
│   │   │   ├── repository.go                      # PortfolioIndexRepository
│   │   │   └── events.go                          # search.portfolio.indexed/reindexed/removed.v1
│   │
│   │   # =========================
│   │   # 🔍 QUERY INPUT & LOGGING
│   │   # =========================
│   │   ├── search_query/
│   │   │   ├── entity.go                          # Query logs (filters/sort/lang/latency, anon user hash)
│   │   │   ├── filters.go                         # Filter structs (rate, skills, languages, location, badges)
│   │   │   ├── errors.go                          # InvalidFilter, BadSort
│   │   │   ├── repository.go                      # SearchQueryRepository
│   │   │   └── events.go                          # search.query.logged / search.query.alert.triggered.v1
│   │   ├── saved_search/                          # ✅ keep existing (explicit saved queries)
│   │   │   ├── entity.go                          # SavedSearch{user_id,name,query_json,schedule,active}
│   │   │   ├── alert.go                           # Alert settings (window, channel)
│   │   │   ├── errors.go                          # SavedSearchNotFound, DuplicateName
│   │   │   ├── repository.go                      # SavedSearchRepository
│   │   │   └── events.go                          # search.saved_search.created/updated/deleted/alert.sent.v1
│   │   ├── multi_language/                        # 🆕 language & analyzer profiles
│   │   │   ├── entity.go                          # DetectedLang{code,confidence}, AnalyzerProfile{id,lang,tokenizer,filters}
│   │   │   ├── detector.go                        # LangDetect(text) → DetectedLang
│   │   │   ├── transliteration.go                 # ar↔en transliteration helpers (names/skills)
│   │   │   ├── repository.go                      # AnalyzerProfileRepository
│   │   │   └── events.go                          # search.lang.profile.updated.v1
│   │   ├── speller/                               # 🆕 spelling & “did you mean”
│   │   │   ├── entity.go                          # SpellingCandidate{term,score,source}
│   │   │   ├── dictionary.go                      # BK-tree / ES suggesters; build/merge
│   │   │   ├── repository.go                      # SpellerRepository
│   │   │   └── events.go                          # search.speller.dictionary.updated.v1
│   │   ├── query_rewrite/                         # 🆕 synonyms/stopwords/rewrites
│   │   │   ├── entity.go                          # RewriteRule{pattern,action,weight,lang,enabled}
│   │   │   ├── rules_engine.go                    # ApplyRules(query, lang) → rewritten query
│   │   │   ├── repository.go                      # RewriteRuleRepository
│   │   │   └── events.go                          # search.query_rewrite.updated.v1
│   │
│   │   # =========================
│   │   # 🧠 RANKING, LTR & CONTROLS
│   │   # =========================
│   │   ├── ltr/                                   # Learning-to-Rank signals
│   │   │   ├── entity.go                          # LTRSignal{doc_id, features, label, effective_at}
│   │   │   ├── signal_source.go                   # Event→feature extractors
│   │   │   ├── feature_store.go                   # Persisted features snapshots
│   │   │   ├── errors.go                          # SignalOutOfRange, FeatureStale
│   │   │   ├── repository.go                      # LTRRepository
│   │   │   └── events.go                          # search.ltr.signal.recorded/features.updated.v1
│   │   ├── promotion/                             # 🆕 editorial boosts/demotions (non-paid)
│   │   │   ├── entity.go                          # Promotion{scope,subject_id,boost,expires_at,reason}
│   │   │   ├── policy.go                          # Guardrails (no paid placements)
│   │   │   ├── repository.go                      # PromotionRepository
│   │   │   └── events.go                          # search.promotion.upserted/expired.v1
│   │
│   │   # =========================
│   │   # 👤 PERSONALIZATION & RECS
│   │   # =========================
│   │   ├── recommendation/
│   │   │   ├── entity.go                          # Recommendation records
│   │   │   ├── score.go                           # Score formula components
│   │   │   ├── reason.go                          # Human-readable reasons
│   │   │   ├── feedback.go                        # User feedback on recs
│   │   │   ├── errors.go                          # ModelUnavailable, ScoringFailed
│   │   │   ├── repository.go                      # RecommendationRepository
│   │   │   └── events.go                          # search.recommendation.generated/feedback.recorded.v1
│   │   ├── recommendation_model/
│   │   │   ├── entity.go                          # ML model metadata/versions
│   │   │   ├── feature.go                         # Feature vectors catalog
│   │   │   ├── training_data.go                   # Training snapshots
│   │   │   ├── errors.go                          # VersionNotFound, FeatureMismatch
│   │   │   ├── repository.go                      # RecommendationModelRepository
│   │   │   └── events.go                          # search.model.version.registered/deployed/rolled_back/training_data.updated.v1
│   │   ├── user_preference/
│   │   │   ├── entity.go                          # Explicit user prefs
│   │   │   ├── implicit_signals.go                # Clicks/views (summaries)
│   │   │   ├── explicit_preferences.go            # Explicit fields (rate, availability)
│   │   │   ├── errors.go                          # PreferenceNotFound
│   │   │   ├── repository.go                      # UserPreferenceRepository
│   │   │   └── events.go                          # search.preference.updated.v1
│   │   ├── personalization/                       # 🆕 profile rollups
│   │   │   ├── entity.go                          # Recent activity, preferred rates, tags
│   │   │   ├── cold_start.go                      # Defaults for new users
│   │   │   ├── errors.go                          # Personalization errors
│   │   │   ├── repository.go                      # PersonalizationRepository
│   │   │   └── events.go                          # search.personalization.cold_start.applied/updated.v1
│   │
│   │   # =========================
│   │   # 🔎 DISCOVERY & SIMILARITY
│   │   # =========================
│   │   ├── matching/
│   │   │   ├── entity.go                          # Job↔Freelancer matches (summary)
│   │   │   ├── criteria.go                        # Criteria (skills, rate, availability)
│   │   │   ├── score_breakdown.go                 # Factor-level scores
│   │   │   ├── errors.go                          # CriteriaInvalid
│   │   │   ├── repository.go                      # MatchingRepository
│   │   │   └── events.go                          # search.match.calculated/accepted/dismissed.v1
│   │   ├── similarity/
│   │   │   ├── entity.go                          # Similar jobs/users (links)
│   │   │   ├── vector.go                          # Vector fields (embeddings)
│   │   │   ├── errors.go                          # Similarity errors
│   │   │   ├── repository.go                      # SimilarityRepository
│   │   │   └── events.go                          # search.similarity.computed/model.updated.v1
│   │   ├── feed/
│   │   │   ├── entity.go                          # User feeds (items, ranks)
│   │   │   ├── item.go                            # Feed item data
│   │   │   ├── personalization.go                 # Feed-specific personalization
│   │   │   ├── errors.go                          # FeedNotFound
│   │   │   ├── repository.go                      # FeedRepository
│   │   │   └── events.go                          # search.feed.item.added/removed/updated.v1
│   │   ├── trending/
│   │   │   ├── entity.go                          # Trending jobs/skills
│   │   │   ├── calculator.go                      # Calculate trending
│   │   │   ├── errors.go                          # Trending errors
│   │   │   ├── repository.go                      # TrendingRepository
│   │   │   └── events.go                          # search.trending.calculated/updated.v1
│   │   ├── suggestion/
│   │   │   └── entity.go                          # Placeholder for suggestion types (lives mostly in application)
│   │
│   │   # =========================
│   │   # 🧭 TAXONOMY & FACETS
│   │   # =========================
│   │   ├── taxonomy/
│   │   │   ├── entity.go                          # Skills/Categories; normalized
│   │   │   ├── category.go                        # Category tree
│   │   │   ├── synonym.go                         # Synonyms & aliases
│   │   │   ├── typos.go                           # Typo tolerance rules
│   │   │   ├── errors.go                          # SkillNotFound, AliasConflict
│   │   │   ├── repository.go                      # TaxonomyRepository
│   │   │   └── events.go                          # search.taxonomy.updated/synonym.changed/typo_rules.updated.v1
│   │   ├── facets/
│   │   │   ├── entity.go                          # Facet defs (price bands, availability, languages, badges)
│   │   │   ├── banding.go                         # Banding logic
│   │   │   ├── tz_overlap.go                      # Timezone overlap
│   │   │   ├── errors.go                          # InvalidBand, UnknownBadge
│   │   │   ├── repository.go                      # FacetRepository
│   │   │   └── events.go                          # search.facets.definition.updated/banding.updated/tz_overlap.updated.v1
│   │
│   │   # =========================
│   │   # 🛡️ SAFETY, HYGIENE & COMPLIANCE
│   │   # =========================
│   │   ├── hygiene/
│   │   │   ├── entity.go                          # Hygiene tasks (incremental, dedupe, archival, visibility)
│   │   │   ├── incremental.go                     # Change markers (version, changed fields)
│   │   │   ├── dedupe.go                          # Fingerprints / duplicates
│   │   │   ├── visibility.go                      # Visibility states
│   │   │   ├── errors.go                          # Hygiene errors
│   │   │   ├── repository.go                      # HygieneRepository
│   │   │   └── events.go                          # search.index.hygiene.*.v1
│   │   ├── compliance/
│   │   │   ├── entity.go                          # Erasure tasks / holds
│   │   │   ├── erasure.go                         # Remove docs on request
│   │   │   └── events.go                          # compliance.erasure.requested/completed.v1
│   │   ├── safety_filters/                        # 🆕 query-time visibility gates
│   │   │   ├── entity.go                          # SafetyRule{kind,subject,action,ttl}
│   │   │   ├── engine.go                          # allow/deny/mask evaluation
│   │   │   └── repository.go                      # SafetyFiltersRepository
│   │
│   │   # =========================
│   │   # 📍 GEO & INTENT
│   │   # =========================
│   │   ├── geo/
│   │   │   ├── entity.go                          # GeoPoint{lat,lon}, GeoPolicy{radius,max_results}
│   │   │   ├── scorer.go                          # Distance decay helpers
│   │   │   └── repository.go                      # GeoRepository
│   │   ├── query_intent/
│   │   │   ├── entity.go                          # Intent{job|talent|navigational, confidence}
│   │   │   ├── classifier.go                      # Heuristics/ML-lite classifier
│   │   │   └── repository.go                      # QueryIntentRepository
│   │
│   │   # =========================
│   │   # 🛠️ OPERATIONS & EXPLAINABILITY
│   │   # =========================
│   │   ├── index_lifecycle/
│   │   │   ├── entity.go                          # IndexSchema{kind,version,alias,mappings_hash}
│   │   │   ├── rollover.go                        # Create v{n+1}, reindex, alias swap
│   │   │   ├── snapshot.go                        # Snapshot/restore via MinIO/local
│   │   │   ├── repository.go                      # IndexLifecycleRepository
│   │   │   └── events.go                          # search.index.rolled_over/snapshotted/restored.v1
│   │   ├── backfill/
│   │   │   ├── entity.go                          # BackfillRun{id,scope,state,counters}
│   │   │   ├── planner.go                         # Partition planning (id/time windows)
│   │   │   └── events.go                          # search.backfill.started/completed.v1
│   │   └── explainability/
│   │       ├── entity.go                          # Explanation{doc_id,factors[],scores[]}
│   │       ├── builder.go                         # Human-readable “why” strings
│   │       └── events.go                          # search.result.explained.v1
│   │
│   # =============================
│   # 📋 APPLICATION LAYER (CQRS)
│   # =============================
│   ├── application/
│   │   # =========================
│   │   # 📡 EVENT CONSUMERS (INBOX)
│   │   # =========================
│   │   ├── eventhandler/
│   │   │   ├── job_handler.go                     # Consumes: job.posted/updated/closed → index jobs & refresh facets
│   │   │   ├── user_handler.go                    # Consumes: user.updated → refresh freelancer docs (badges, visibility)
│   │   │   ├── review_handler.go                  # Consumes: review.* → update rating aggregates, LTR signals
│   │   │   ├── entitlement_handler.go             # Consumes: subscription.feature.changed → gate facets/features
│   │   │   ├── admin_content_handler.go           # Consumes: admin.content.actioned → hide/unhide docs
│   │   │   ├── admin_flags_handler.go             # Consumes: admin.feature_flag/threshold/experiment.updated → refresh toggles
│   │   │   ├── storage_lifecycle_handler.go       # Consumes: file.lifecycle.soft_deleted/restored → sync asset visibility in docs
│   │   │   ├── compliance_handler.go              # Consumes: user.erasure.requested → remove indexed docs
│   │   │   └── taxonomy_handler.go                # Consumes: admin.taxonomy.updated → refresh synonyms/rewrites
│   │
│   │   # =========================
│   │   # 🧠 USE CASES (COMMANDS/QUERIES)
│   │   # =========================
│   │   # ---- 🔍 SEARCH EXECUTION ----
│   │   ├── search/
│   │   │   ├── service.go                         # Run job/talent searches end-to-end
│   │   │   ├── job_search.go                      # ES DSL for jobs
│   │   │   ├── freelancer_search.go               # ES DSL for users
│   │   │   ├── query_builder.go                   # Query + filter composition
│   │   │   ├── facet_builder.go                   # Aggregations/facets builder
│   │   │   ├── dto.go                             # Request/Response DTOs
│   │   │   ├── mapper.go                          # Map ES hits → DTO
│   │   │   ├── commands.go                        # ExecuteSearch, SaveSearch
│   │   │   ├── queries.go                         # GetSearchResults, GetSearchSuggestions
│   │   │   └── validators.go                      # Validate filters/sorts/facets
│   │   # ---- 🧾 INDEXING ----
│   │   ├── indexing/
│   │   │   ├── service.go                         # Index/update/remove docs
│   │   │   ├── job_indexer.go                     # Job document mappers
│   │   │   ├── user_indexer.go                    # User document mappers
│   │   │   ├── bulk_indexer.go                    # Bulk ops & backpressure
│   │   │   ├── dto.go                             # Indexing DTOs
│   │   │   ├── commands.go                        # IndexJob, IndexUser, ReindexAll
│   │   │   ├── queries.go                         # GetIndexStatus, GetDocByID
│   │   │   └── validators.go                      # Payload validation
│   │   # ---- 💾 SAVED SEARCHES ----
│   │   ├── saved_search/
│   │   │   ├── service.go                         # Create/Update/Delete saved searches + alerts
│   │   │   ├── commands.go                        # CreateSavedSearch, UpdateSavedSearch, DeleteSavedSearch
│   │   │   ├── queries.go                         # GetSavedSearches, GetSavedSearch
│   │   │   ├── validators.go                      # Name uniqueness, schedule bounds
│   │   │   ├── dto.go                             
│   │   │   └── mapper.go                          
│   │   # ---- 👤 PERSONALIZATION & RECS ----
│   │   ├── recommendation/
│   │   │   ├── service.go                         # Generate & explain recs
│   │   │   ├── job_recommender.go                 # Jobs → freelancer
│   │   │   ├── freelancer_recommender.go          # Freelancers → client
│   │   │   ├── collaborative_filtering.go         # CF signals
│   │   │   ├── content_based.go                   # Content-based signals
│   │   │   ├── hybrid_recommender.go              # Hybrid approach
│   │   │   ├── scoring_engine.go                  # Score fusion
│   │   │   ├── personalization.go                 # Per-user personalization
│   │   │   ├── diversity_optimizer.go             # Result diversity
│   │   │   ├── cold_start_handler.go              # New users/jobs
│   │   │   ├── dto.go
│   │   │   ├── ml_model.go                        # Model selection
│   │   │   ├── commands.go                        # GenerateRecommendations, RecordFeedback
│   │   │   ├── queries.go                         # GetRecommendations, GetReasons
│   │   │   └── validators.go                      # Limits & params
│   │   ├── user_preference/
│   │   │   ├── service.go                         # Update/Get prefs
│   │   │   ├── commands.go                        # UpdatePreferences
│   │   │   └── queries.go                         # GetPreferences
│   │   ├── personalization/
│   │   │   ├── service.go                         # UpdateProfile, ComputeDefaults
│   │   │   ├── commands.go                        # UpdatePersonalizationProfile
│   │   │   └── queries.go                         # GetPersonalizationProfile
│   │   # ---- 🔎 DISCOVERY ----
│   │   ├── matching/
│   │   │   ├── service.go                         # Orchestrate criteria evaluators
│   │   │   ├── matcher.go                         # Match pipeline
│   │   │   ├── criteria_evaluator.go              # Evaluate match criteria
│   │   │   ├── skill_matcher.go                   # Skill overlap
│   │   │   ├── experience_matcher.go              # Experience fit
│   │   │   ├── rate_matcher.go                    # Rate fit
│   │   │   ├── availability_matcher.go            # Availability/tz fit
│   │   │   ├── score_calculator.go                # Overall score
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                        # CreateMatchRun
│   │   │   ├── queries.go                         # GetMatchesForJob/User
│   │   │   └── validators.go                      # Criteria completeness
│   │   ├── similarity/
│   │   │   ├── service.go                         # Similar jobs/users
│   │   │   ├── job_similarity.go
│   │   │   ├── user_similarity.go
│   │   │   ├── vector_calculator.go               # Embedding generation
│   │   │   ├── dto.go
│   │   │   ├── commands.go                        # RebuildSimilarityVectors
│   │   │   ├── queries.go                         # GetSimilarJobs/Users
│   │   │   └── validators.go                      # k/threshold
│   │   ├── feed/
│   │   │   ├── service.go                         # Generate personalized feed
│   │   │   ├── generator.go                       # Candidate gather
│   │   │   ├── ranking.go                         # Ranker
│   │   │   ├── freshness_scorer.go                # Freshness factor
│   │   │   ├── relevance_scorer.go                # Relevance factor
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                        # GenerateFeed
│   │   │   ├── queries.go                         # GetFeed
│   │   │   └── validators.go                      # Window/size checks
│   │   ├── trending/
│   │   │   ├── service.go                         # Trending recompute
│   │   │   ├── calculator.go
│   │   │   ├── dto.go
│   │   │   ├── mapper.go
│   │   │   ├── commands.go                        # RecomputeTrending
│   │   │   ├── queries.go                         # GetTrending
│   │   │   └── validators.go                      # Window/minSupport
│   │   ├── suggestion/
│   │   │   ├── service.go                         # Autocomplete suggestions
│   │   │   ├── dto.go
│   │   │   ├── cache_warmer.go                    # Warm cache
│   │   │   ├── commands.go                        # WarmSuggestionCache
│   │   │   ├── queries.go                         # GetSuggestions
│   │   │   └── validators.go                      # Prefix/lang checks
│   │   ├── portfolio_index/
│   │   │   ├── service.go                         # Index & search portfolios
│   │   │   ├── commands.go                        # IndexPortfolioDoc
│   │   │   └── queries.go                         # SearchPortfolios
│   │   # ---- 🧭 TAXONOMY & FACETS ----
│   │   ├── taxonomy/
│   │   │   ├── service.go                         # Upsert skills/categories/aliases
│   │   │   ├── commands.go                        # UpsertSkill/Category, Add/RemoveAlias
│   │   │   ├── queries.go                         # GetSkill, ListSkills, GetCategoryTree
│   │   │   ├── validators.go                      # Alias conflicts, edit distance
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── facets/
│   │   │   ├── service.go                         # Build facets; apply banding
│   │   │   ├── commands.go                        # DefineFacet, UpdateFacetBands
│   │   │   ├── queries.go                         # GetFacetsForQuery
│   │   │   ├── validators.go                      # Band & tz rules
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   # ---- 🛡️ SAFETY & HYGIENE ----
│   │   ├── hygiene/
│   │   │   ├── service.go                         # Incrementals, dedupe, archive/visibility
│   │   │   ├── commands.go                        # RunIncrementalUpdate, RunDedup, ArchiveDoc, SetVisibility
│   │   │   ├── queries.go                         # GetHygieneStatus, GetDocHistory
│   │   │   ├── validators.go                      # State transitions
│   │   │   ├── dto.go
│   │   │   └── mapper.go
│   │   ├── compliance/
│   │   │   ├── service.go                         # Process erasure
│   │   │   ├── commands.go                        # RequestErasure
│   │   │   └── queries.go                         # GetErasureStatus
│   │   ├── safety_filters/
│   │   │   ├── service.go                         # ApplySafetyRules to result sets
│   │   │   ├── commands.go                        # UpsertSafetyRule, RemoveSafetyRule
│   │   │   └── queries.go                         # ListSafetyRules
│   │   # ---- 📍 GEO & INTENT ----
│   │   ├── geo/
│   │   │   ├── service.go                         # ApplyGeoFilters, ScoreByDistance
│   │   │   └── validators.go                      # Radius/bounds checks
│   │   ├── query_intent/
│   │   │   ├── service.go                         # ClassifyIntent (job vs talent vs navigational)
│   │   │   └── queries.go                         # GetIntent
│   │   # ---- 🛠️ OPERATIONS ----
│   │   ├── index_lifecycle/
│   │   │   ├── service.go                         # Rollover, Snapshot, Restore
│   │   │   ├── commands.go                        # RolloverIndex, SnapshotIndex, RestoreIndex
│   │   │   └── queries.go                         # GetIndexSchema
│   │   ├── backfill/
│   │   │   ├── service.go                         # Plan & run backfills
│   │   │   └── commands.go                        # StartBackfill
│   │   └── explainability/
│   │       ├── service.go                         # BuildExplanation (ES _explain wrappers)
│   │       └── queries.go                         # GetExplanation
│   │
│   # =============================
│   # 🔌 INFRASTRUCTURE LAYER
│   # =============================
│   ├── infrastructure/
│   │   # =========================
│   │   # 🗄️ PERSISTENCE (POSTGRES)
│   │   # =========================
│   │   ├── persistence/
│   │   │   └── postgres/
│   │   │       # 🧱 COMMON (DB BOOTSTRAP)
│   │   │       ├── connection.go                   # DSN & pooling
│   │   │       ├── transaction.go                  # TX helpers
│   │   │       ├── migrations.go                   # Auto-migrate (versioned)
│   │   │       ├── version.go                      # Schema version table
│   │   │       ├── safety.go                       # Pre-flight checks (env/disk)
│   │   │       # 🔍 QUERY INPUT & LOGGING
│   │   │       ├── search_query_repository.go      # SearchQueryRepository impl
│   │   │       ├── saved_search_repository.go      # SavedSearchRepository impl
│   │   │       ├── speller_repository.go           # 🆕 SpellerRepository impl
│   │   │       ├── query_rewrite_repository.go     # 🆕 RewriteRuleRepository impl
│   │   │       ├── multi_language_repository.go    # 🆕 AnalyzerProfileRepository impl
│   │   │       # 👤 PERSONALIZATION & RECS
│   │   │       ├── recommendation_repository.go
│   │   │       ├── recommendation_model_repository.go
│   │   │       ├── user_preference_repository.go
│   │   │       ├── personalization_repository.go   # 🆕
│   │   │       ├── ltr_repository.go               # 🆕 LTRRepository impl
│   │   │       # 🔎 DISCOVERY
│   │   │       ├── matching_repository.go
│   │   │       ├── similarity_repository.go
│   │   │       ├── feed_repository.go
│   │   │       ├── trending_repository.go
│   │   │       ├── portfolio_repository.go         # 🆕 portfolios index meta (PG)
│   │   │       # 🧭 TAXONOMY & FACETS
│   │   │       ├── taxonomy_repository.go          # 🆕
│   │   │       ├── facets_repository.go            # 🆕
│   │   │       # 🛡️ SAFETY/HYGIENE/COMPLIANCE
│   │   │       ├── hygiene_repository.go           # 🆕
│   │   │       ├── compliance_repository.go        # 🆕
│   │   │       ├── safety_filters_repository.go    # 🆕
│   │   │       # 📍 GEO & INTENT
│   │   │       ├── geo_repository.go               # 🆕
│   │   │       └── query_intent_repository.go      # 🆕
│   │   # =========================
│   │   # 🔎 ELASTICSEARCH
│   │   # =========================
│   │   ├── elasticsearch/
│   │   │   ├── client.go                          # ES client (connection, retry)
│   │   │   ├── index_manager.go                   # Create/update mappings/settings
│   │   │   ├── alias_router.go                    # Read/write aliases (single-tenant)
│   │   │   ├── snapshot_client.go                 # Snapshot/restore to MinIO/local
│   │   │   ├── job_mapper.go                      # Job entity → ES document
│   │   │   ├── user_mapper.go                     # User entity → ES document
│   │   │   └── config.go                          # Hosts, auth, timeouts
│   │   # =========================
│   │   # ⚡ CACHE (REDIS)
│   │   # =========================
│   │   ├── cache/
│   │   │   └── redis/
│   │   │       ├── connection.go                  # Pooling, retries, instrumentation
│   │   │       ├── search_cache.go                # Cache search results (TTL)
│   │   │       ├── suggestion_cache.go            # Autocomplete cache
│   │   │       ├── feed_cache.go                  # User feeds
│   │   │       ├── recommendation_cache.go        # Recommendations
│   │   │       ├── taxonomy_cache.go              # Skills/categories & aliases
│   │   │       ├── ltr_cache.go                   # LTR feature snapshots
│   │   │       ├── speller_cache.go               # Did-you-mean cache
│   │   │       ├── rewrite_rules_cache.go         # Query rewrite rules
│   │   │       └── promotion_cache.go             # Active promotions
│   │   # =========================
│   │   # 📨 MESSAGING (KAFKA)
│   │   # =========================
│   │   ├── messaging/
│   │   │   └── kafka/
│   │   │       ├── consumer.go                    # 📝 uses platform-shared/inbox for dedupe & offsets
│   │   │       ├── producer.go                    # 📝 uses platform-shared/outbox for reliable publishing
│   │   │       ├── topics.go                      # 📝 contracts/events: job.*, user.*, review.*, experiment.*, compliance.*
│   │   │       └── scram.go                       # SASL/SCRAM-256 auth
│   │   # =========================
│   │   # 🤖 ML UTILITIES
│   │   # =========================
│   │   └── ml/
│   │       ├── model_loader.go                    # Load models (disk/env)
│   │       ├── predictor.go                       # Predictions/rerank
│   │       ├── feature_extractor.go               # Extract features
│   │       ├── trainer.go                         # Offline training jobs
│   │       └── evaluator.go                       # Evaluate model performance
│   │
│   # =============================
│   # 🌐 HTTP INTERFACE (v1)
│   # =============================
│   ├── interfaces/
│   │   └── http/
│   │       └── v1/
│   │           # =========================
│   │           # 🧭 HANDLERS
│   │           # =========================
│   │           ├── handlers/
│   │           # ---- SEARCH & INDEXING ----
│   │           │   ├── search_handler.go           # POST /search/jobs, /search/freelancers
│   │           │   ├── indexing_handler.go         # POST/GET admin index ops
│   │           │   ├── saved_search_handler.go     # 🆕 CRUD saved searches & alerts
│   │           # ---- RECS / DISCOVERY ----
│   │           │   ├── recommendation_handler.go   # GET /recommendations/*
│   │           │   ├── matching_handler.go         # POST/GET /matching/*
│   │           │   ├── feed_handler.go             # GET /feed/*
│   │           │   ├── trending_handler.go         # GET /trending/*
│   │           │   ├── similarity_handler.go       # GET /similarity/*
│   │           │   ├── suggestion_handler.go       # GET /suggestions/*
│   │           │   ├── portfolio_handler.go        # 🆕 GET /search/portfolios
│   │           # ---- TAXONOMY & FACETS ----
│   │           │   ├── taxonomy_handler.go         # GET/PUT /taxonomy/*
│   │           │   ├── facets_handler.go           # GET /facets/*
│   │           # ---- INPUT HELPERS ----
│   │           │   ├── speller_handler.go          # 🆕 GET /speller/did-you-mean
│   │           │   ├── rewrite_handler.go          # 🆕 POST /rewrites/preview
│   │           │   ├── language_handler.go         # 🆕 GET /languages/supported
│   │           # ---- SAFETY & OPS ----
│   │           │   ├── hygiene_handler.go          # POST/GET /hygiene/*
│   │           │   ├── compliance_handler.go       # POST /compliance/erasure
│   │           │   ├── lifecycle_handler.go        # POST /indices/rollover|snapshot|restore
│   │           │   ├── promotion_handler.go        # PUT/DELETE /promotions/*
│   │           │   ├── explain_handler.go          # GET /debug/explain
│   │           │   └── health_handler.go           # GET /health, /ready, /live
│   │           # =========================
│   │           # 🧰 MIDDLEWARE
│   │           # =========================
│   │           ├── middleware/
│   │           │   ├── auth.go                     # JWT verification (pkg/auth)
│   │           │   ├── rbac.go                     # Role checks
│   │           │   ├── cors.go                     # CORS (platform-shared/ginx)
│   │           │   ├── rate_limit.go               # Token bucket rate limiting
│   │           │   ├── logging.go                  # Structured request logging
│   │           │   ├── error_handler.go            # Error → HTTP mapping
│   │           │   └── request_id.go               # X-Request-ID propagation
│   │           # =========================
│   │           # 📨 RESPONSES
│   │           # =========================
│   │           ├── responses/
│   │           │   ├── success.go                  # platform-shared/httpx/response
│   │           │   ├── error.go                    # platform-shared/httpx/errors
│   │           │   └── pagination.go               # platform-shared/httpx/pagination
│   │           # =========================
│   │           # 🗺️ ROUTES (SECTIONED)
│   │           # =========================
│   │           ├── routes/
│   │           │   # ---- SEARCH & INDEXING ----
│   │           │   ├── search_routes.go            # /search/* (jobs|freelancers)
│   │           │   ├── indexing_routes.go          # /indexing/* (admin)
│   │           │   ├── saved_search_routes.go      # 🆕 /saved-searches/*
│   │           │   # ---- RECS / DISCOVERY ----
│   │           │   ├── recommendation_routes.go    # /recommendations/*
│   │           │   ├── matching_routes.go          # /matching/*
│   │           │   ├── feed_routes.go              # /feed/*
│   │           │   ├── trending_routes.go          # /trending/*
│   │           │   ├── similarity_routes.go        # /similarity/*
│   │           │   ├── suggestion_routes.go        # /suggestions/*
│   │           │   ├── portfolio_routes.go         # 🆕 /search/portfolios/*
│   │           │   # ---- TAXONOMY & FACETS ----
│   │           │   ├── taxonomy_routes.go          # /taxonomy/*
│   │           │   ├── facets_routes.go            # /facets/*
│   │           │   # ---- INPUT HELPERS ----
│   │           │   ├── speller_routes.go           # 🆕 /speller/*
│   │           │   ├── rewrite_routes.go           # 🆕 /rewrites/*
│   │           │   ├── language_routes.go          # 🆕 /languages/*
│   │           │   # ---- SAFETY & OPS ----
│   │           │   ├── hygiene_routes.go           # /hygiene/*
│   │           │   ├── compliance_routes.go        # /compliance/*
│   │           │   ├── lifecycle_routes.go         # /indices/*
│   │           │   ├── promotion_routes.go         # /promotions/*
│   │           │   └── websocket_routes.go         # (optional) /ws if needed later
│   │           └── router.go                       # Gin engine wiring + common middleware
│   │
│   └── (end internal)
│
├── config/
│   ├── default.yaml                               # Defaults
│   ├── dev.yaml                                   # Dev overrides
│   └── prod.yaml                                  # Prod overrides
│
├── dapr/
│   ├── local/
│   │   ├── pubsub.yaml                            # Kafka pub/sub
│   │   └── statestore.yaml                        # State store
│   └── k8s/
│       ├── pubsub.yaml                            # Scopes: ["search-be"]
│       ├── statestore.yaml                        # Scopes
│       └── secrets.yaml                           # Secret store
│
├── elasticsearch/
│   # =============================
│   # 🧭 ES ARTIFACTS (STATIC)
│   # =============================
│   ├── mappings/
│   │   ├── jobs.json                               # Job index mapping
│   │   ├── users.json                              # User index mapping
│   │   └── portfolios.json                         # 🆕 Portfolio index mapping
│   └── analyzers/
│       └── custom_analyzers.json                   # ICU folding, ar/en analyzers, n-grams
│
├── ml_models/
│   ├── job_recommendation/
│   │   ├── model.pkl
│   │   ├── features.json
│   │   └── metadata.json
│   ├── freelancer_recommendation/
│   │   ├── model.pkl
│   │   ├── features.json
│   │   └── metadata.json
│   └── matching/
│       ├── model.pkl
│       ├── features.json
│       └── metadata.json
│
├── deployments/
│   └── k8s/
│       ├── deployment.yaml                        # Deployment
│       ├── service.yaml                           # Service
│       ├── configmap.yaml                         # ConfigMap
│       ├── secrets.yaml                           # Secrets
│       ├── hpa.yaml                               # HPA
│       ├── pdb.yaml                               # PDB
│       └── servicemonitor.yaml                    # Prometheus ServiceMonitor
│
├── scripts/
│   ├── setup-local.sh                             # Local environment bootstrap
│   ├── get-secrets.sh                             # Fetch secrets
│   ├── seed-data.sh                               # Seed PG + ES for dev
│   ├── create-indices.sh                          # Create ES indices & aliases
│   ├── reindex-all.sh                             # Rollover & reindex pipeline
│   └── train-models.sh                            # Offline training
│
├── tests/
│   # =============================
│   # ✅ TEST SUITES
│   # =============================
│   ├── unit/
│   │   ├── domain/                                # Rewrite/speller/facets/taxonomy/LTR unit tests
│   │   ├── application/                           # Search/indexing/recs services tests
│   │   └── infrastructure/                        # Repos, ES mappers tests
│   ├── integration/
│   │   ├── handlers/                              # HTTP integration tests
│   │   └── repositories/                          # Postgres repositories tests
│   └── e2e/
│       └── scenarios/                             # Search→click→LTR signal flows
│
├── docs/
│   ├── README.md                                  # Service overview
│   ├── api.md                                     # API reference
│   ├── events.md                                  # Published/Consumed events (contracts/events)
│   ├── search-algorithms.md                       # Ranking & filtering details
│   ├── recommendation-engine.md                   # Recommenders & signals
│   ├── recommendation-types.md                    # Placement taxonomy
│   ├── matching-algorithm.md                      # Matching pipeline
│   ├── ml-models.md                               # Models & features
│   ├── elasticsearch-setup.md                     # ES setup & snapshots
│   ├── MIGRATIONS.md                              # Migration history
│   ├── SCHEMA.md                                  # Database schema
│   └── RUNBOOK.md                                 # Ops procedures (reindex, hygiene, snapshots, backfills)
│
├── pkg/
│   # =============================
│   # 🧰 LOCAL UTILITIES
│   # =============================
│   ├── errors/
│   │   ├── errors.go                               # Service-specific error helpers
│   │   └── codes.go                                # Error codes
│   ├── utils/
│   │   ├── validator.go                            # Validation utilities
│   │   ├── text_analyzer.go                        # Tokenize/normalize helpers
│   │   ├── normalizer.go                           # Data normalization
│   │   └── vector_math.go                          # Vector ops & distances
│   └── constants/
│       ├── indices.go                               # Index/alias names
│       └── README.md                                # (Note) events/topics come from contracts/events
│
├── .github/
│   └── workflows/
│       ├── ci.yml                                  # CI pipeline
│       └── cd.yml                                  # CD pipeline
│
├── go.mod                                          # 📝 imports pkg/auth, platform-shared, contracts/events
├── go.sum
├── .env.example
├── Makefile
├── Dockerfile
├── .dockerignore
├── .gitignore
└── README.md


```
