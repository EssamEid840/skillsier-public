### Skillsier Database DesignConventions

*   **Primary keys:** id UUIDv7 (time-ordered).
    
*   **Timestamps:** created\_at, updated\_at, deleted\_at (soft delete).
    
*   **PII encryption:** email, names, phone, addresses, DOB, location, postal\_code, optional lat/lon, URLs/proof files.
    
*   **Case-insensitive uniques:** lower(username), lower(email).
    
*   **Sharding:** hash by user.id for hot/high-volume tables (Citus-ready).
    
*   **Partitioning (monthly):** audit\_logs, risk\_signal, client\_job\_analytics, user\_statistics\_history.
    
*   **Events:** every CUD → outbox\_events (\*.v1 topics).
    
*   **Idempotency:** HTTP (http\_idempotency) + consumer (inbox\_messages).
    
*   **Enums:** superset of enums.go; when synonyms differ (e.g., unresponsive vs unresponsive\_communication) keep unresponsive\_communication and alias in domain logic.
    

Domain: user (core)
===================

user
----

*   id PK
    
*   keycloak\_id varchar(255) unique
    
*   username varchar(50) unique on lower(username)
    
*   email varchar(255) unique on lower(email) encrypted
    
*   first\_name, middle\_name, last\_name varchar(100) encrypted
    
*   full\_name varchar(200) // computed for search
    
*   display\_name varchar(200) NULL
    
*   user\_type enum(freelancer, client, org, admin, moderator, support)
    
*   additional\_types jsonb NULL // hybrid roles
    
*   status enum(pending, active, inactive, suspended, banned, deleted, restricted)
    
*   verification\_status enum(unverified, pending, verified, rejected, expired, require\_resubmit)
    
*   email\_verified, phone\_verified, identity\_verified bool
    
*   last\_login\_at, last\_seen\_at timestamptz
    
*   date\_of\_birth date encrypted
    
*   gender enum from enums.go NULL
    
*   nationality varchar(100) NULL
    
*   languages text\[\] NULL
    
*   profile\_picture\_url, cover\_image\_url, video\_intro\_url varchar(255) NULL
    
*   thumbnail\_url varchar(255) NULL
    
*   bio text encrypted NULL
    
*   tagline varchar(200) NULL
    
*   title varchar(200) NULL
    
*   overview text encrypted NULL
    
*   website varchar(500) NULL
    
*   social\_links jsonb NULL // \[{platform,url}\]
    
*   rating numeric(3,2) NULL
    
*   total\_reviews int default 0
    
*   completed\_jobs int default 0
    
*   total\_jobs int default 0
    
*   success\_rate numeric(5,2) NULL
    
*   total\_earnings numeric(15,2) NULL
    
*   total\_spent numeric(15,2) NULL
    
*   response\_time\_avg interval NULL
    
*   availability\_status enum(available, busy, unavailable, part\_time, full\_time, on\_leave)
    
*   hours\_per\_week int NULL
    
*   accepting\_work bool default false
    
*   badges text\[\] NULL
    
*   is\_featured, is\_top\_rated, is\_rising\_talent, is\_expert\_vetted bool default false
    
*   profile\_completeness int default 0
    
*   profile\_completed bool default false
    
*   version int // optimistic concurrency
    
*   data\_residency enum(EU, US) default US
    
*   legal\_hold bool default false // prevents hard-deletes
    
*   export\_in\_progress bool default false
    
*   locale varchar(10) NULL
    
*   content\_region varchar(2) NULL
    
*   onboarding\_stage enum(created, profile\_basic, profile\_complete, verified, ready)
    
*   onboarding\_checklist jsonb NULL
    
*   acquisition\_channel varchar(64) NULL
    
*   acquisition\_source varchar(64) NULL
    
*   acquisition\_campaign varchar(128) NULL
    
*   first\_referrer\_url varchar(500) NULL
    
*   ab\_bucket varchar(64) NULL
    
*   feature\_flags jsonb NULL // read-only projection
    
*   moderation\_status enum(clean, limited, shadow, suspended)
    
*   shadow\_reasons text\[\] NULL
    
*   spam\_score numeric(5,2) NULL
    
*   trust\_tier enum(low, normal, elevated, gold)
    
*   communication\_opt\_outs jsonb NULL
    
*   consent\_version varchar(32) NULL
    
*   consent\_accepted\_at timestamptz NULL
    
*   cookie\_consent jsonb NULL
    
*   data\_export\_requested\_at timestamptz NULL
    
*   last\_data\_export\_delivered\_at timestamptz NULL
    
*   last\_terms\_ack\_at timestamptz NULL
    
*   last\_privacy\_ack\_at timestamptz NULL
    
*   last\_mfa\_challenge\_at timestamptz NULL
    
*   failed\_login\_streak int default 0
    
*   marketing\_tags text\[\] NULL
    
*   support\_tier enum(standard, priority, enterprise)

*   linked_professional_accounts (jsonb NULL)

*   misclassification_risk (enum: low, medium, high NULL)
    
*   **Personalization/trust extras:**personalization\_score numeric(5,2) NULLvendor\_rating numeric(3,2) NULLgig\_preferences jsonb NULLreview\_summary text NULL
    
*   created\_at, updated\_at, deleted\_at
    

**Constraints & rules**

*   DOB validity: 1900-01-01 < date\_of\_birth < now if provided.
    
*   Hybrid roles: if biz uses synthetic “both”, ensure additional\_types includes freelancer and client.
    
*   Emit user.onboarding.updated.v1 when stage/checklist changes.
    
*   Keep user.data\_residency aligned with settings.data\_residency.region (emit migration event).
    
*   Activity sort by last\_login\_at DESC; expose languages, badges to search-be.
    

user\_phone
-----------

*   id PK
    
*   user\_id FK → user.id
    
*   e164 varchar(32) encrypted
    
*   is\_primary bool default false
    
*   verified\_at timestamptz NULL
    
*   created\_at, updated\_at, deleted\_atRules: one primary per user (partial unique); cap 5 phones/user.
    

user\_address
-------------

*   id PK
    
*   user\_id FK → user.id
    
*   label varchar(50) NULL
    
*   country varchar(2)
    
*   region varchar(100) encrypted NULL
    
*   city varchar(100) encrypted
    
*   postal\_code varchar(20) encrypted NULL
    
*   line1, line2 varchar(255) encrypted NULL
    
*   is\_primary bool default false
    
*   geo\_hash varchar(12) NULL
    
*   created\_at, updated\_at, deleted\_atRules: one primary per user; cap 5/user.
    

user\_email\_alias
------------------

*   id PK
    
*   user\_id FK → user.id
    
*   email varchar(255) encrypted
    
*   is\_primary bool default false
    
*   verified\_at timestamptz NULL
    
*   created\_at, updated\_at, deleted\_atRules: one primary (and must be verified); cap 5/user.
    

username\_history
-----------------

*   id PK
    
*   user\_id FK → user.id
    
*   old\_username, new\_username varchar(50)
    
*   changed\_at timestamptz
    
*   reason varchar(255) NULL
    
*   approved\_by UUID NULL
    
*   created\_at, updated\_at, deleted\_at
    

user\_referral
--------------

*   id PK
    
*   user\_id FK → user.id unique
    
*   referrer\_id FK → user.id NULL
    
*   referral\_code varchar(20) unique
    
*   bonus\_connects\_granted int default 0
    
*   applied\_at timestamptz
    
*   reward\_tier enum(basic, premium) NULL
    
*   created\_at, updated\_at, deleted\_at
    

user\_statistics (1:1)
----------------------

*   id PK
    
*   user\_id FK → user.id unique
    
*   last\_30d\_logins int
    
*   messages\_sent\_30d int
    
*   proposals\_submitted\_30d int
    
*   contracts\_active int
    
*   disputes\_count\_12m int
    
*   engagement\_score numeric(5,2) NULL
    
*   ltv\_prediction numeric(15,2) NULL
    
*   cohort\_id varchar(32) NULL
    
*   created\_at, updated\_at, deleted\_at
    

user\_statistics\_history (partitioned)
---------------------------------------

*   id PK
    
*   user\_id FK → user.id
    
*   period\_start date
    
*   period\_end date
    
*   last\_30d\_logins, messages\_sent\_30d, proposals\_submitted\_30d, contracts\_active, disputes\_count\_12m int
    
*   engagement\_score numeric(5,2) NULL
    
*   ltv\_prediction numeric(15,2) NULL
    
*   churn\_risk numeric(5,2) NULL
    
*   created\_at, updated\_at, deleted\_atPartition by period\_start monthly.
    

user\_warning
-------------

*   id PK
    
*   user\_id FK → user.id
    
*   reason enum(late\_delivery, poor\_quality, unresponsive\_communication, minor\_violation)
    
*   severity enum(low, medium, high, critical)
    
*   message text NULL
    
*   issued\_at timestamptz
    
*   issued\_by UUID
    
*   acknowledged\_at timestamptz NULL
    
*   escalation\_level int default 0
    
*   created\_at, updated\_at, deleted\_at
    

user\_device\_locale
--------------------

*   id PK
    
*   user\_id FK → user.id
    
*   device\_locale varchar(10)
    
*   first\_seen\_at, last\_seen\_at timestamptz
    
*   created\_at, updated\_at, deleted\_at
    

user\_consent
-------------

*   id PK
    
*   user\_id FK → user.id
    
*   policy\_slug varchar(64)
    
*   version varchar(32)
    
*   accepted bool
    
*   accepted\_at timestamptz
    
*   revoked\_at timestamptz NULL
    
*   meta jsonb // ip, ua, geo snapshot

*   audit_trail (jsonb NULL)
    
*   created\_at, updated\_at, deleted\_at
    

user\_review\_summary
---------------------

*   id PK
    
*   user\_id FK → user.id unique
    
*   positive\_themes text\[\] NULL
    
*   negative\_themes text\[\] NULL
    
*   ai\_generated\_at timestamptz
    
*   created\_at, updated\_at, deleted\_at
    

Domain: profile
===============

profile (1:1 user)
------------------

*   id PK
    
*   user\_id FK → user.id unique
    
*   bio text encrypted
    
*   tagline varchar(255)
    
*   title varchar(100)
    
*   location varchar(255) encrypted
    
*   timezone varchar(50)
    
*   completion\_percentage int
    
*   profile\_picture\_url varchar(255) NULL
    
*   preferences\_json jsonb
    
*   visibility enum(public, invited, private, anonymous) default public
    
*   version int
    
*   headline\_keywords text\[\] NULL
    
*   public\_metrics jsonb NULL
    
*   seo\_slug varchar(200) NULL
    
*   primary\_category varchar(64) NULL
    
*   secondary\_categories text\[\] NULL
    
*   language\_tags text\[\] NULL
    
*   client\_focus\_keywords text\[\] NULL
    
*   certification\_badges text\[\] NULL

*   video_specs (jsonb NULL)
    
*   created\_at, updated\_at, deleted\_at
    

profile\_preferences (1:1)
--------------------------

*   id PK
    
*   profile\_id FK → profile.id unique
    
*   language varchar(10) default 'en'
    
*   currency varchar(3) default 'USD'
    
*   work\_type enum(hourly, fixed, both)
    
*   min\_rate, max\_rate numeric(10,2) NULL
    
*   preferred\_contract\_types text\[\] NULL
    
*   remote\_preference enum(remote, hybrid, onsite, any)
    
*   preferred\_industries text\[\] NULL
    
*   exclude\_industries text\[\] NULL
    
*   gig\_duration\_prefs enum(short, medium, long, ongoing)
    
*   created\_at, updated\_at, deleted\_at
    

profile\_availability (1:1)
---------------------------

*   id PK
    
*   profile\_id FK → profile.id unique
    
*   availability\_status enum(available, busy, unavailable, part\_time, full\_time, on\_leave)
    
*   hours\_per\_week int NULL
    
*   next\_available\_at timestamptz NULL
    
*   availability\_note varchar(255) NULL
    
*   timezone\_offset\_minutes int2 NULL
    
*   preferred\_shift enum(morning, afternoon, evening, night) NULL
    
*   created\_at, updated\_at, deleted\_at
    

profile\_locale\_variant (1:N)
------------------------------

*   id PK
    
*   profile\_id FK → profile.id
    
*   locale varchar(10)
    
*   translated\_bio text encrypted
    
*   translated\_tagline varchar(255)
    
*   translated\_title varchar(100)
    
*   translated\_overview text encrypted
    
*   created\_at, updated\_at, deleted\_atUnique: (profile\_id, locale)
    

profile\_snapshot (1:N)
-----------------------

*   id PK
    
*   profile\_id FK → profile.id
    
*   snapshot\_data jsonb
    
*   snapshot\_reason varchar(255)
    
*   restored\_at timestamptz NULL
    
*   created\_at, updated\_at, deleted\_at
    

profile\_geo (1:1)
------------------

*   id PK
    
*   profile\_id FK → profile.id unique
    
*   country varchar(2)
    
*   region varchar(100) encrypted NULL
    
*   city varchar(100) encrypted NULL
    
*   lat, lon numeric(9,6) encrypted NULL
    
*   geo\_confidence enum(low, medium, high)
    
*   preferred\_radius\_km int NULL
    
*   created\_at, updated\_at, deleted\_at
    

Domain: skill
=============

skill
-----

*   id PK
    
*   name varchar(100) unique
    
*   taxonomy\_id UUID NULL
    
*   popularity\_rank int NULL
    
*   created\_at, updated\_at, deleted\_at
    

user\_skill (junction; PK(user\_id, skill\_id))
-----------------------------------------------

*   user\_id FK → user.id
    
*   skill\_id FK → skill.id
    
*   proficiency enum(beginner, intermediate, advanced, expert)
    
*   years\_experience int
    
*   display\_order int
    
*   last\_used\_at timestamptz NULL
    
*   relevance\_score numeric(5,2) NULL
    
*   evidence\_count int default 0

*   endorsement_count (int default 0)
    
*   certified\_by varchar(64) NULL
    
*   created\_at, updated\_at, deleted\_at
    

skill\_endorsement
------------------

*   id PK
    
*   skill\_id FK → skill.id
    
*   endorser\_id FK → user.id
    
*   endorsed\_user\_id FK → user.id
    
*   endorsement\_date timestamptz
    
*   evidence\_url varchar(255) NULL
    
*   weight numeric(4,2) default 1.0
    
*   endorser\_trust\_score numeric(5,2) NULL
    
*   created\_at, updated\_at, deleted\_at
    

skill\_test
-----------

*   id PK
    
*   skill\_id FK → skill.id
    
*   user\_id FK → user.id
    
*   score numeric(5,2)
    
*   test\_date timestamptz
    
*   percentile\_rank numeric(5,2) NULL
    
*   created\_at, updated\_at, deleted\_at
    

skill\_pricing\_tier
--------------------

*   id PK
    
*   skill\_id FK → skill.id
    
*   tier enum(beginner, intermediate, expert)
    
*   hourly\_rate numeric(10,2)
    
*   min\_projects\_required int default 0
    
*   created\_at, updated\_at, deleted\_atUnique: (skill\_id, tier)
    

Domain: experience
==================

experience
----------

*   id PK
    
*   user\_id FK → user.id
    
*   company, title varchar(100)
    
*   description text
    
*   start\_date date
    
*   end\_date date NULL
    
*   is\_current bool default false
    
*   location\_country varchar(2) NULL
    
*   skills\_highlight text\[\] NULL
    
*   outcome\_metrics jsonb NULL
    
*   budget\_range jsonb NULL // {min,max,currency}

*   nda_signed (bool default false)  # Tracks NDAs; from compliance research (e.g., data confidentiality in freelancing).
    
*   created\_at, updated\_at, deleted\_at
    

experience\_reference (1:N)
---------------------------

*   id PK
    
*   experience\_id FK → experience.id
    
*   reference\_text text
    
*   reference\_url varchar(255) NULL
    
*   client\_name varchar(100) NULL
    
*   rating numeric(3,2) NULL
    
*   verified bool default false
    
*   contact\_consent bool default false
    
*   created\_at, updated\_at, deleted\_at
    

experience\_verification (1:1)
------------------------------

*   id PK
    
*   experience\_id FK → experience.id unique
    
*   status enum(pending, verified, rejected)
    
*   verified\_at timestamptz NULL
    
*   rejection\_reason varchar(255) NULL
    
*   proof\_url varchar(255) NULL
    
*   created\_at, updated\_at, deleted\_at
    

experience\_gap (1:N)
---------------------

*   id PK
    
*   user\_id FK → user.id
    
*   start\_date, end\_date date
    
*   gap\_explanation text NULL
    
*   flagged bool default false
    
*   gap\_type enum(career\_break, education, travel) NULL
    
*   created\_at, updated\_at, deleted\_at
    

Domain: education
=================

education
---------

*   id PK
    
*   user\_id FK → user.id
    
*   school, degree, field varchar(100)
    
*   graduation\_year int
    
*   description text NULL
    
*   accreditation\_level enum(diploma, bachelor, master, phd, other) NULL
    
*   gpa numeric(3,2) NULL
    
*   honors text\[\] NULL
    
*   created\_at, updated\_at, deleted\_at
    

education\_link (1:N)
---------------------

*   id PK
    
*   education\_id FK → education.id
    
*   url varchar(255)
    
*   verified bool default false
    
*   link\_type enum(alumni, transcript) NULL
    
*   created\_at, updated\_at, deleted\_at
    

education\_attachment (1:N)
---------------------------

*   id PK
    
*   education\_id FK → education.id
    
*   file\_url varchar(255) encrypted
    
*   file\_type enum(transcript, gpa, certificate)
    
*   metadata jsonb
    
*   created\_at, updated\_at, deleted\_at
    

Domain: certification
=====================

certification
-------------

*   id PK
    
*   user\_id FK → user.id
    
*   name, issuing\_organization varchar(100)
    
*   issue\_date date
    
*   expiry\_date date NULL
    
*   credential\_id varchar(50) NULL
    
*   url varchar(255) NULL
    
*   status enum(pending, verified, rejected, expired)
    
*   issuer\_url varchar(255) NULL
    
*   verification\_url varchar(255) NULL
    
*   scope text\[\] NULL
    
*   renewal\_required bool default false
    
*   created\_at, updated\_at, deleted\_at
    

certification\_proctoring (1:1)
-------------------------------

*   id PK
    
*   certification\_id FK → certification.id unique
    
*   metadata jsonb
    
*   proctor\_score numeric(5,2) NULL
    
*   created\_at, updated\_at, deleted\_at
    

Domain: portfolio
=================

portfolio\_item
---------------

*   id PK
    
*   user\_id FK → user.id
    
*   title varchar(100)
    
*   description text
    
*   url varchar(255) NULL
    
*   thumbnail\_url varchar(255) NULL
    
*   display\_order int default 0
    
*   category enum(design, development, writing, other)
    
*   impact\_score numeric(5,2) NULL
    
*   team\_size int NULL
    
*   duration\_weeks int NULL
    
*   client\_industry varchar(100) NULL
    
*   views\_count int default 0
    
*   created\_at, updated\_at, deleted\_at
    

portfolio\_item\_detail (1:1)
-----------------------------

*   id PK
    
*   portfolio\_item\_id FK → portfolio\_item.id unique
    
*   tags text\[\]/jsonb
    
*   tech\_stack text\[\]/jsonb
    
*   role varchar(100) NULL
    
*   outcome text NULL
    
*   budget numeric(12,2) NULL
    
*   currency varchar(3) NULL
    
*   created\_at, updated\_at, deleted\_at
    

portfolio\_media (1:N)
----------------------

*   id PK
    
*   portfolio\_item\_id FK → portfolio\_item.id
    
*   file\_url varchar(255)
    
*   file\_type enum(image, video, document)
    
*   processed bool default false
    
*   width, height, duration\_seconds int NULL
    
*   alt\_text varchar(255) NULL
    
*   created\_at, updated\_at, deleted\_at
    

portfolio\_share\_link (1:N)
----------------------------

*   id PK
    
*   portfolio\_item\_id FK → portfolio\_item.id
    
*   link varchar(255) unique
    
*   expires\_at timestamptz NULL
    
*   views\_count int default 0
    
*   created\_at, updated\_at, deleted\_at
    

Domain: language
================

language
--------

*   id PK
    
*   user\_id FK → user.id
    
*   code varchar(10)
    
*   proficiency enum(basic, conversational, fluent, native)
    
*   cefr\_level enum(A1, A2, B1, B2, C1, C2) NULL
    
*   use\_contexts text\[\] NULL
    
*   verified\_by varchar(64) NULL
    
*   accent varchar(64) NULL
    
*   created\_at, updated\_at, deleted\_atUnique: (user\_id, code)
    

language\_certificate (1:N)
---------------------------

*   id PK
    
*   language\_id FK → language.id
    
*   certificate\_url varchar(255)
    
*   verified bool default false
    
*   created\_at, updated\_at, deleted\_at
    

language\_fluency\_test (1:N)
-----------------------------

*   id PK
    
*   language\_id FK → language.id
    
*   score numeric(5,2)
    
*   test\_date timestamptz
    
*   test\_provider varchar(64) NULL
    
*   created\_at, updated\_at, deleted\_at
    

Domain: freelancer
==================

freelancer (1:1 user)
---------------------

*   id PK
    
*   user\_id FK → user.id unique
    
*   title varchar(100) NULL
    
*   overview text NULL
    
*   video\_intro\_url varchar(255) NULL
    
*   jss\_score numeric(5,2) NULL
    
*   success\_rate numeric(5,2) NULL
    
*   total\_jobs int
    
*   total\_earnings numeric(15,2)
    
*   response\_time\_avg interval NULL
    
*   tier enum(basic, pro, elite)
    
*   vetting\_status enum(not\_applied, applied, in\_review, approved, rejected)
    
*   niche\_tags text\[\] NULL
    
*   service\_packages jsonb NULL
    
*   preferred\_tools text\[\] NULL
    
*   travel\_availability enum(none, regional, global)
    
*   search\_boost numeric(4,2) NULL

*   pro_eligibility_score (numeric(5,2) NULL)  # Score for Pro status; from Fiverr Pro applications, with detailed verification.
    
*   response\_rate\_90d numeric(5,2) NULL
    
*   created\_at, updated\_at, deleted\_at
    

freelancer\_profile (1:1)
-------------------------

*   id PK
    
*   freelancer\_id FK → freelancer.id unique
    
*   tagline varchar(255)
    
*   service\_offering text
    
*   unique\_selling\_points text\[\] NULL
    
*   created\_at, updated\_at, deleted\_at
    

freelancer\_rates (1:1)
-----------------------

*   id PK
    
*   freelancer\_id FK → freelancer.id unique
    
*   hourly\_rate numeric(10,2) NULL
    
*   minimum\_budget numeric(12,2) NULL
    
*   currency varchar(3)
    
*   discount\_tier jsonb NULL
    
*   created\_at, updated\_at, deleted\_at
    

freelancer\_stats (1:1)
-----------------------

*   id PK
    
*   freelancer\_id FK → freelancer.id unique
    
*   total\_jobs int
    
*   total\_earnings numeric(15,2)
    
*   success\_rate numeric(5,2)
    
*   response\_time\_avg interval NULL
    
*   jss\_score numeric(5,2)
    
*   tier enum(basic, pro, elite)
    
*   vetting\_status enum(not\_applied, applied, in\_review, approved, rejected)
    
*   repeat\_clients\_rate numeric(5,2) NULL
    
*   avg\_project\_value numeric(12,2) NULL
    
*   proposals\_win\_rate\_90d numeric(5,2) NULL
    
*   response\_time\_p50\_ms int NULL
    
*   response\_time\_p95\_ms int NULL
    
*   client\_satisfaction\_score numeric(5,2) NULL
    
*   created\_at, updated\_at, deleted\_at
    

freelancer\_rate\_band (1:N)
----------------------------

*   id PK
    
*   freelancer\_id FK → freelancer.id
    
*   band\_type enum(hourly, daily, retainer)
    
*   rate numeric(10,2)
    
*   created\_at, updated\_at, deleted\_at
    

freelancer\_connect (1:1)
-------------------------

*   id PK
    
*   freelancer\_id FK → freelancer.id unique
    
*   balance int default 0
    
*   auto\_top\_up\_threshold int NULL
    
*   auto\_top\_up\_amount int NULL
    
*   lifetime\_spent int default 0
    
*   created\_at, updated\_at, deleted\_at
    

freelancer\_subscription\_bundle (1:N)
--------------------------------------

*   id PK
    
*   freelancer\_id FK → freelancer.id
    
*   bundle\_name varchar(50)
    
*   features jsonb
    
*   price numeric(10,2)
    
*   subscribed\_at timestamptz
    
*   renewal\_date timestamptz NULL
    
*   created\_at, updated\_at, deleted\_at
    

Domain: client
==============

client (1:1 user)
-----------------

*   id PK
    
*   user\_id FK → user.id unique
    
*   company\_name varchar(100) NULL
    
*   company\_size enum(small, medium, large) NULL
    
*   industry varchar(100) NULL
    
*   total\_hires int default 0
    
*   total\_spent numeric(15,2) default 0
    
*   active\_contracts int default 0
    
*   average\_rating numeric(3,2) NULL
    
*   tier enum(basic, premium)
    
*   spend\_tier enum(new, bronze, silver, gold, platinum)
    
*   payment\_verification\_level enum(none, basic, advanced)
    
*   billing\_country varchar(2) NULL
    
*   hiring\_categories text\[\] NULL
    
*   preferred\_freelancer\_tiers text\[\] NULL
    
*   created\_at, updated\_at, deleted\_at
    

client\_profile (1:1)
---------------------

*   id PK
    
*   client\_id FK → client.id unique
    
*   about text NULL
    
*   website varchar(255) NULL
    
*   location varchar(255) NULL
    
*   team\_size int NULL

*   vendor_management_level (enum: basic, advanced NULL) # Enterprise vendor tracking; from ERD freelance schemas.
    
*   created\_at, updated\_at, deleted\_at
    

client\_company (1:1)
---------------------

*   id PK
    
*   client\_id FK → client.id unique
    
*   name varchar(100)
    
*   size enum(small, medium, large)
    
*   industry varchar(100)
    
*   founded\_year int NULL
    
*   employees int NULL
    
*   linkedin\_url varchar(255) NULL
    
*   crunchbase\_url varchar(255) NULL
    
*   revenue\_band enum(<1M, "1–10M", "10–100M", "100M+")
    
*   created\_at, updated\_at, deleted\_at
    

client\_stats (1:1)
-------------------

*   id PK
    
*   client\_id FK → client.id unique
    
*   total\_hires int
    
*   total\_spent numeric(15,2)
    
*   active\_contracts int
    
*   average\_rating numeric(3,2)
    
*   tier enum(basic, premium)
    
*   retention\_rate numeric(5,2) NULL
    
*   created\_at, updated\_at, deleted\_at
    

client\_payment\_method (1:N)
-----------------------------

*   id PK
    
*   client\_id FK → client.id
    
*   method\_type enum(credit\_card, paypal, bank\_transfer)
    
*   details jsonb encrypted
    
*   verified bool default false
    
*   last\_used\_at timestamptz NULL
    
*   created\_at, updated\_at, deleted\_at
    

client\_job\_analytics (1:N, partitioned)
-----------------------------------------

*   id PK
    
*   client\_id FK → client.id
    
*   hire\_rate numeric(5,2)
    
*   response\_time\_avg interval
    
*   period\_start date
    
*   period\_end date
    
*   budget\_utilization numeric(5,2) NULL
    
*   created\_at, updated\_at, deleted\_atPartitioned monthly by period\_start.
    

Domain: verification (KYC/KYB)
==============================

verification
------------

*   id PK
    
*   entity\_id FK → user.id or org.id
    
*   entity\_type enum(user, org)
    
*   type enum(email, phone, identity, social, mobile, kyb)
    
*   status enum(pending, verified, rejected, expired)
    
*   submitted\_at timestamptz
    
*   verified\_at timestamptz NULL
    
*   rejection\_reason varchar(255) NULL
    
*   tier enum(light, standard, enhanced)
    
*   provider varchar(64) NULL
    
*   provider\_reference varchar(128) NULL
    
*   attempts int default 0
    
*   cooldown\_until timestamptz NULL
    
*   risk\_category enum(low, medium, high)
    
*   flags text\[\] NULL
    
*   biometric\_match\_score numeric(5,2) NULL

*   manual_review_notes (text NULL) # Admin notes; from Fiverr regulatory verification.
    
*   created\_at, updated\_at, deleted\_at
    

verification\_document (1:N)
----------------------------

*   id PK
    
*   verification\_id FK → verification.id
    
*   doc\_type enum(id\_card, passport, proof\_of\_address, selfie, other)
    
*   file\_url varchar(255) encrypted
    
*   doc\_hash varchar(64) NULL
    
*   metadata jsonb
    
*   review\_status enum(pending, accepted, rejected) default pending
    
*   ocr\_extracted jsonb NULL
    
*   created\_at, updated\_at, deleted\_at
    

verification\_audit (1:N)
-------------------------

*   id PK
    
*   verification\_id FK → verification.id
    
*   action enum(submitted, approved, rejected, reverified)
    
*   actor\_id UUID NULL
    
*   notes text NULL
    
*   created\_at timestamptz
    
*   ip\_logged inet NULL
    
*   created\_at, updated\_at, deleted\_at
    

sanctions\_screening (1:1)
--------------------------

*   id PK
    
*   verification\_id FK → verification.id unique
    
*   screened\_at timestamptz
    
*   flagged bool
    
*   details jsonb
    
*   last\_screened\_at timestamptz NULL
    
*   created\_at, updated\_at, deleted\_at
    

international\_tax (1:1)
------------------------

*   id PK
    
*   verification\_id FK → verification.id unique
    
*   vat\_gst\_number varchar(50)
    
*   validated\_at timestamptz
    
*   valid bool
    
*   validation\_source varchar(64) NULL
    
*   last\_validation\_at timestamptz NULL
    
*   created\_at, updated\_at, deleted\_at
    

Domain: settings
================

settings (1:1 user)
-------------------

*   id PK
    
*   user\_id FK → user.id unique
    
*   theme enum(light, dark) default light
    
*   language varchar(10) default 'en'
    
*   timezone varchar(50) default 'UTC'
    
*   currency varchar(3) default 'USD'
    
*   default\_view enum(dashboard, jobs, messages) NULL
    
*   created\_at, updated\_at, deleted\_at
    

### notification\_pref (1:1)

*   id PK
    
*   settings\_id FK → settings.id unique
    
*   default\_email bool default true
    
*   default\_sms bool default false
    
*   default\_push bool default true
    
*   default\_in\_app bool default true
    
*   channels\_matrix jsonb
    
*   transport\_tokens jsonb NULL
    
*   digest\_prefs jsonb NULL
    
*   priority\_topics text\[\] NULL
    
*   created\_at, updated\_at, deleted\_at
    

### privacy\_settings (1:1)

*   id PK
    
*   settings\_id FK → settings.id unique
    
*   profile\_visibility enum(public, invited, private, anonymous) default public
    
*   show\_email bool default false
    
*   show\_phone bool default false
    
*   searchable\_profile bool default true
    
*   hide\_rate\_history bool default false
    
*   hide\_online\_status bool default false
    
*   block\_invitations\_from text\[\] NULL
    
*   data\_sharing\_consent bool default false
    
*   created\_at, updated\_at, deleted\_at
    

### data\_residency (1:1)

*   id PK
    
*   settings\_id FK → settings.id unique
    
*   region enum(EU, US) default US
    
*   migration\_status enum(pending, completed, failed) default completed
    
*   preferred\_cloud\_provider varchar(32) NULL
    
*   created\_at, updated\_at, deleted\_at
    

### accessibility (1:1)

*   id PK
    
*   settings\_id FK → settings.id unique
    
*   high\_contrast bool default false
    
*   screen\_reader bool default false
    
*   font\_size\_pref enum(small, medium, large)
    
*   created\_at, updated\_at, deleted\_at
    

Domain: saved\_items
====================

saved\_item
-----------

*   id PK
    
*   user\_id FK → user.id
    
*   item\_type enum(job, freelancer)
    
*   item\_id UUID
    
*   notes text NULL
    
*   saved\_at timestamptz
    
*   expiry\_at timestamptz NULL
    
*   pinned bool default false
    
*   labels text\[\] NULL
    
*   notification\_on\_update bool default false
    
*   created\_at, updated\_at, deleted\_atUnique: (user\_id, item\_type, item\_id)
    

saved\_collection
-----------------

*   id PK
    
*   user\_id FK → user.id
    
*   name varchar(100)
    
*   description text NULL
    
*   created\_at, updated\_at, deleted\_at
    

collection\_item (junction)
---------------------------

*   collection\_id FK → saved\_collection.id
    
*   saved\_item\_id FK → saved\_item.id
    
*   added\_at timestamptzPK: (collection\_id, saved\_item\_id)
    

collection\_share
-----------------

*   id PK
    
*   collection\_id FK → saved\_collection.id
    
*   shared\_with\_user\_id FK → user.id
    
*   permissions enum(view, edit)
    
*   expires\_at timestamptz NULL
    
*   message varchar(255) NULL
    
*   access\_logs jsonb NULL
    
*   created\_at, updated\_at, deleted\_at
    

Domain: blocked\_users
======================

blocked\_user
-------------

*   id PK
    
*   blocker\_id FK → user.id
    
*   blocked\_id FK → user.id
    
*   reason varchar(255) NULL
    
*   scope enum(messaging, invites, full) default full
    
*   duration interval NULL
    
*   blocked\_at timestamptz
    
*   expires\_at timestamptz NULL
    
*   origin enum(user\_action, auto\_moderation, admin\_action)
    
*   evidence jsonb NULL
    
*   mutual\_block bool default false
    
*   created\_at, updated\_at, deleted\_atUnique: (blocker\_id, blocked\_id); CHECK: blocker\_id <> blocked\_id
    

block\_appeal (1:N)
-------------------

*   id PK
    
*   blocked\_user\_id FK → blocked\_user.id
    
*   appeal\_text text
    
*   appeal\_reason text NULL
    
*   status enum(pending, approved, rejected)
    
*   reviewed\_at timestamptz NULL
    
*   resolution\_notes text NULL
    
*   evidence\_attachments text\[\] NULL
    
*   created\_at, updated\_at, deleted\_at
    

Domain: user\_suspension
========================

user\_suspension
----------------

*   id PK
    
*   user\_id FK → user.id
    
*   reason enum(tos\_violation, payment\_issue, quality\_issues, abusive\_behavior)
    
*   start\_date timestamptz
    
*   end\_date timestamptz NULL
    
*   suspended\_by UUID NULL
    
*   is\_active bool default true
    
*   auto\_detector varchar(64) NULL
    
*   review\_due\_at timestamptz NULL
    
*   appeal\_eligible bool default true
    
*   created\_at, updated\_at, deleted\_at
    

suspension\_history (1:N)
-------------------------

*   id PK
    
*   suspension\_id FK → user\_suspension.id
    
*   action enum(placed, extended, released)
    
*   notes text NULL
    
*   action\_at timestamptz
    
*   created\_at, updated\_at, deleted\_at
    

Domain: user\_ban
=================

user\_ban
---------

*   id PK
    
*   user\_id FK → user.id
    
*   reason enum(fraud, severe\_abuse, multiple\_violations, security\_threat)
    
*   banned\_at timestamptz
    
*   banned\_by UUID NULL
    
*   is\_permanent bool default true
    
*   expires\_at timestamptz NULL
    
*   linked\_accounts text\[\] NULL
    
*   public\_message varchar(255) NULL
    
*   ban\_source varchar(64) NULL
    
*   created\_at, updated\_at, deleted\_at
    

ban\_history (1:N)
------------------

*   id PK
    
*   ban\_id FK → user\_ban.id
    
*   action enum(placed, extended, released)
    
*   notes text NULL
    
*   action\_at timestamptz
    
*   created\_at, updated\_at, deleted\_at
    

ban\_evasion (1:N)
------------------

*   id PK
    
*   ban\_id FK → user\_ban.id
    
*   detected\_user\_id FK → user.id NULL ON DELETE SET NULL
    
*   signals jsonb
    
*   detected\_at timestamptz
    
*   confidence\_score numeric(5,2) NULL
    
*   created\_at, updated\_at, deleted\_at
    

Domain: org
===========

org
---

*   id PK
    
*   name varchar(100) unique
    
*   owner\_id FK → user.id
    
*   billing\_profile\_id UUID NULL
    
*   seat\_limit int default 5
    
*   parent\_org\_id FK → org.id NULL
    
*   plan\_slug varchar(64) NULL
    
*   data\_residency enum(EU, US)
    
*   tags text\[\] NULL
    
*   employee\_count int NULL

*   metadata jsonb NULL

*   compliance_level (enum: basic, certified NULL) # Reason: Enterprise compliance certification; from research.
    
*   created\_at, updated\_at, deleted\_at
    

org\_member (junction; PK(org\_id, user\_id))
---------------------------------------------

*   org\_id FK → org.id
    
*   user\_id FK → user.id
    
*   role enum(owner, admin, member)
    
*   permissions jsonb NULL
    
*   invited\_by UUID NULL
    
*   invitation\_accepted\_at timestamptz NULL
    
*   joined\_at timestamptz
    
*   last\_active\_at timestamptz NULL
    
*   created\_at, updated\_at, deleted\_at
    

org\_seat (1:N)
---------------

*   id PK
    
*   org\_id FK → org.id
    
*   user\_id FK → user.id NULL // null = unassigned
    
*   status enum(assigned, unassigned, pending\_invite)
    
*   cost\_center varchar(64) NULL
    
*   created\_at, updated\_at, deleted\_at
    

org\_talent\_pool (1:N)
-----------------------

*   id PK
    
*   org\_id FK → org.id
    
*   name varchar(100)
    
*   pool\_type enum(internal, external) NULL
    
*   created\_at, updated\_at, deleted\_at
    

talent\_pool\_member (junction; PK(talent\_pool\_id, user\_id))
---------------------------------------------------------------

*   talent\_pool\_id FK → org\_talent\_pool.id
    
*   user\_id FK → user.id
    
*   added\_at timestamptz
    

org\_invite (1:N)
-----------------

*   id PK
    
*   org\_id FK → org.id
    
*   email varchar(255) encrypted
    
*   role enum(owner, admin, member)
    
*   invite\_token\_hash bytea
    
*   expires\_at timestamptz
    
*   accepted\_at timestamptz NULL
    
*   invite\_message text NULL
    
*   created\_at, updated\_at, deleted\_at
    

Domain: security\_center
========================

security\_settings (1:1 user)
-----------------------------

*   id PK
    
*   user\_id FK → user.id unique
    
*   two\_fa\_enabled bool default false
    
*   recovery\_keys\_hash bytea encrypted NULL
    
*   mfa\_methods text\[\] NULL // sms, totp, webauthn, backup\_codes
    
*   passwordless\_enabled bool default false
    
*   sso\_provider varchar(64) NULL
    
*   biometric\_enabled bool default false
    
*   created\_at, updated\_at, deleted\_at
    

device (1:N)
------------

*   id PK
    
*   user\_id FK → user.id
    
*   fingerprint varchar(255)
    
*   platform varchar(32) NULL
    
*   app\_version varchar(32) NULL
    
*   os\_version varchar(32) NULL
    
*   last\_seen\_at timestamptz
    
*   revoked bool default false
    
*   risk enum(low, medium, high) NULL
    
*   created\_at, updated\_at, deleted\_atUnique (recommended): (user\_id, fingerprint)
    

session (1:N)
-------------

*   id PK
    
*   user\_id FK → user.id
    
*   ip inet
    
*   user\_agent varchar(255)
    
*   geo\_ip\_country varchar(2) NULL
    
*   ip\_asn int NULL
    
*   client\_fingerprint\_hash bytea NULL
    
*   device\_id FK → device.id NULL
    
*   created\_at timestamptz
    
*   expires\_at timestamptz
    
*   revoked bool default false
    
*   updated\_at timestamptz
    
*   deleted\_at timestamptz
    

passkey (1:N)
-------------

*   id PK
    
*   user\_id FK → user.id
    
*   credential bytea encrypted
    
*   last\_used\_at timestamptz NULL
    
*   created\_at, updated\_at, deleted\_at
    

login\_alert (1:N)
------------------

*   id PK
    
*   user\_id FK → user.id
    
*   session\_id FK → session.id NULL ON DELETE SET NULL
    
*   alert\_type enum(new\_device, suspicious\_location)
    
*   severity enum(info, warn, critical)
    
*   channel\_used text\[\] NULL
    
*   sent\_at timestamptz
    
*   resolved bool default false
    
*   created\_at, updated\_at, deleted\_at
    

Domain: compliance
==================

tax\_profile (1:1 user)
-----------------------

*   id PK
    
*   user\_id FK → user.id unique
    
*   country varchar(50)
    
*   vat\_gst varchar(50) NULL
    
*   tin varchar(50) encrypted NULL
    
*   w\_form\_url varchar(255) NULL
    
*   validated bool default false
    
*   validation\_source varchar(64) NULL
    
*   last\_validation\_at timestamptz NULL
    
*   created\_at, updated\_at, deleted\_at
    

residency (1:1 user)
--------------------

*   id PK
    
*   user\_id FK → user.id unique
    
*   country varchar(50)
    
*   since date
    
*   proof\_url varchar(255) NULL
    
*   review\_status enum(pending, verified, rejected)
    
*   reviewed\_at timestamptz NULL
    
*   created\_at, updated\_at, deleted\_at
    

compliance\_artifact (1:N)
--------------------------

*   id PK
    
*   user\_id FK → user.id
    
*   type enum(w9, vat, other)
    
*   file\_url varchar(255) encrypted
    
*   expiry\_date date NULL
    
*   created\_at, updated\_at, deleted\_at
    

audit\_log\_retention (1:1 user or global)
------------------------------------------

*   id PK
    
*   user\_id FK → user.id NULL
    
*   retention\_period interval
    
*   override\_reason varchar(255) NULL
    
*   created\_at, updated\_at, deleted\_at
    

breach\_notification (1:N)
--------------------------

*   id PK
    
*   user\_id FK → user.id
    
*   sent\_at timestamptz
    
*   details jsonb
    
*   acknowledged\_at timestamptz NULL
    
*   created\_at, updated\_at, deleted\_at
    

user\_legal\_hold (1:N)
-----------------------

*   id PK
    
*   user\_id FK → user.id
    
*   reason varchar(255)
    
*   placed\_by UUID NULL
    
*   released\_at timestamptz NULL
    
*   case\_id varchar(64) NULL
    
*   created\_at, updated\_at, deleted\_at
    

data\_export\_job (1:N)
-----------------------

*   id PK
    
*   user\_id FK → user.id
    
*   status enum(queued, running, delivered, failed)
    
*   storage\_url varchar(255) NULL
    
*   checksum varchar(64) NULL
    
*   requested\_at timestamptz
    
*   completed\_at timestamptz NULL
    
*   format enum(zip, pdf) NULL
    
*   created\_at, updated\_at, deleted\_at
    

Domain: risk\_signals
=====================

risk\_signal (partition monthly by occurred\_at)
------------------------------------------------

*   id PK
    
*   user\_id FK → user.id
    
*   type enum(ip\_mismatch, dispute, chargeback, velocity\_alert)
    
*   severity enum(low, medium, high)
    
*   occurred\_at timestamptz
    
*   details jsonb
    
*   ip\_country varchar(2) NULL
    
*   device\_graph\_id varchar(64) NULL
    
*   velocity\_window varchar(16) NULL
    
*   related\_entities jsonb NULL
    
*   mitigation\_action varchar(64) NULL
    
*   created\_at, updated\_at, deleted\_at
    

risk\_score (1:1)
-----------------

*   id PK
    
*   user\_id FK → user.id unique
    
*   score numeric(5,2)
    
*   model\_version varchar(32) NULL
    
*   reason\_codes text\[\] NULL
    
*   updated\_at timestamptz
    
*   created\_at, deleted\_at
    

risk\_hold (1:N)
----------------

*   id PK
    
*   user\_id FK → user.id
    
*   type enum(account, payment)
    
*   reason varchar(255)
    
*   actor varchar(64) NULL
    
*   severity enum(low, medium, high)
    
*   until timestamptz NULL
    
*   notified\_at timestamptz NULL
    
*   created\_at, updated\_at, deleted\_at
    

risk\_alert (1:N)
-----------------

*   id PK
    
*   user\_id FK → user.id
    
*   alert\_type enum(high\_risk\_action, anomaly)
    
*   sent\_at timestamptz
    
*   details jsonb
    
*   escalated\_to UUID NULL
    
*   created\_at, updated\_at, deleted\_at
    

Domain: profile\_depth
======================

profile\_depth (1:1 profile)
----------------------------

*   id PK
    
*   profile\_id FK → profile.id unique
    
*   normalized\_skills jsonb
    
*   badges jsonb // \[{slug, awarded\_at, reason}\]
    
*   badge\_sources jsonb NULL
    
*   depth\_score numeric(5,2)
    
*   marketplace\_fit\_score numeric(5,2) NULL
    
*   skill\_coverage numeric(5,2) NULL
    
*   created\_at, updated\_at, deleted\_at
    

profile\_rate\_history (1:N)
----------------------------

*   id PK
    
*   profile\_id FK → profile.id
    
*   amount numeric(10,2)
    
*   currency varchar(3)
    
*   effective\_at timestamptz
    
*   rationale varchar(255) NULL
    
*   created\_at, updated\_at, deleted\_atUnique: (profile\_id, effective\_at)
    

profile\_availability\_slot (1:N)
---------------------------------

*   id PK
    
*   profile\_id FK → profile.id
    
*   weekday int2
    
*   start\_minutes int2
    
*   end\_minutes int2
    
*   tz varchar(50)
    
*   is\_recurring bool default true
    
*   until\_date date NULL
    
*   slot\_type enum(meeting, work) NULL
    
*   created\_at, updated\_at, deleted\_atUnique: (profile\_id, weekday, start\_minutes, end\_minutes, tz)
    

profile\_taxonomy\_map (1:N)
----------------------------

*   id PK
    
*   profile\_id FK → profile.id
    
*   source\_skill varchar(100)
    
*   normalized\_skill varchar(100)
    
*   confidence numeric(4,3)
    
*   source varchar(64) NULL
    
*   created\_at, updated\_at, deleted\_atUnique: (profile\_id, source\_skill)
    

Domain: cross-cutting
=====================

audit\_logs (partition monthly)
-------------------------------

*   id PK
    
*   entity\_type varchar(50)
    
*   entity\_id UUID
    
*   action varchar(50)
    
*   actor\_id UUID NULL // user/admin/system
    
*   changes jsonb
    
*   ip inet
    
*   user\_agent varchar(255)
    
*   request\_id varchar(64) NULL
    
*   region varchar(2) NULL
    
*   correlation\_id varchar(64) NULL
    
*   created\_at timestamptz
    
*   updated\_at, deleted\_at
    

outbox\_events
--------------

*   id PK
    
*   aggregate\_id UUID
    
*   topic varchar(200)
    
*   payload jsonb
    
*   status enum(pending, published, failed)
    
*   attempt int default 0
    
*   available\_at timestamptz
    
*   trace jsonb NULL
    
*   created\_at, updated\_at, deleted\_at
    

inbox\_messages
---------------

*   idempotency\_key varchar(128) PK
    
*   handler varchar(100)
    
*   processed\_at timestamptz NULL
    
*   retry\_count int default 0
    
*   created\_at, updated\_at, deleted\_at
    

http\_idempotency
-----------------

*   key varchar(128) PK
    
*   response\_status int
    
*   response\_body bytea
    
*   expires\_at timestamptz
    
*   request\_fingerprint bytea NULL
    
*   created\_at, updated\_at, deleted\_at
    

projection\_search\_cache (read-optimized)
------------------------------------------

*   id PK
    
*   user\_id FK → user.id unique
    
*   profile\_id FK → profile.id unique
    
*   tokens text\[\]
    
*   facets jsonb // country, rate\_band, categories, langs, badges…
    
*   freshness timestamptz
    
*   last\_synced\_at timestamptz
    
*   created\_at, updated\_at, deleted\_at
    

Relations (summary)
===================

*   **User core:** user 1:1 → profile, settings, security\_settings, user\_statistics, risk\_score, tax\_profile, residency, profile\_depth, freelancer?, client?, user\_review\_summary.user 1:N → user\_phone, user\_address, user\_email\_alias, username\_history, saved\_item, user\_warning, user\_suspension, user\_ban, risk\_signal, risk\_hold, device, session, login\_alert, compliance\_artifact, verification, experience, education, certification, portfolio\_item, language, org\_member, user\_statistics\_history, user\_consent, user\_device\_locale.user ↔ user M:N via blocked\_user; referrals via user\_referral(referrer\_id).
    
*   **Profile:** 1:1 → profile\_preferences, profile\_availability, profile\_geo, profile\_depth; 1:N → profile\_locale\_variant, profile\_snapshot, profile\_rate\_history, profile\_availability\_slot, profile\_taxonomy\_map.
    
*   **Skills:** user M:N via user\_skill; skill 1:N → skill\_endorsement, skill\_test, skill\_pricing\_tier.
    
*   **Freelancer/Client:** user 1:1 → freelancer / client; satellites as defined.
    
*   **Verification:** user/org 1:N verification; verification 1:N verification\_document, verification\_audit; 1:1 sanctions\_screening, international\_tax.
    
*   **Security:** user 1:N device/session/passkey/login\_alert; session.device\_id → device.id.
    
*   **Compliance/Risk:** user 1:N risk\_signal/risk\_hold/risk\_alert/compliance\_artifact/breach\_notification/user\_legal\_hold/data\_export\_job.
    
*   **Org:** org M:N user via org\_member; org 1:N org\_seat/org\_invite/org\_talent\_pool; pools M:N user via talent\_pool\_member; org hierarchy via parent\_org\_id.
    
*   **Cross-cutting:** audit\_logs.actor\_id → user.id (nullable).
    

Indexing & Ops (concise checklist; no SQL)
==========================================

**Search & GIN**

*   Full-text on titles/descriptions: user, profile, experience, education, portfolio\_item.
    
*   GIN on arrays/jsonb: badges, language\_tags, profile.tokens/facets, portfolio\_item\_detail.tags/tech\_stack, client\_focus\_keywords, certification\_badges, unique\_selling\_points, priority\_topics, positive\_themes, negative\_themes, gig\_preferences.
    

**Geospatial**

*   Btree on user\_address.geo\_hash for coarse filtering; precise lat/lon kept encrypted.
    

**Hot composites**

*   user: (user\_type, status, verification\_status), (last\_login\_at DESC), (personalization\_score DESC), (vendor\_rating DESC)
    
*   profile: (visibility), (work\_type, min\_rate, max\_rate) (via preferences), (availability\_status, hours\_per\_week), (next\_available\_at)
    
*   skill: (popularity\_rank ASC NULLS LAST); user\_skill (user\_id, display\_order)
    
*   experience: (user\_id, start\_date DESC), (user\_id, is\_current)
    
*   education: (user\_id, graduation\_year DESC)
    
*   certification: (user\_id, status), (expiry\_date)
    
*   portfolio: (user\_id, display\_order), (category), portfolio\_item (views\_count DESC)
    
*   language: (user\_id, proficiency DESC)
    
*   freelancer: (tier, jss\_score DESC), (vetting\_status), (response\_rate\_90d DESC); freelancer\_stats (client\_satisfaction\_score DESC)
    
*   client: (tier, average\_rating DESC); client\_stats (retention\_rate DESC)
    
*   client\_job\_analytics: (period\_start DESC), (budget\_utilization DESC)
    
*   verification: (entity\_type, entity\_id, type), (tier), (biometric\_match\_score DESC); international\_tax (valid)
    
*   saved\_item: (user\_id, item\_type), (expiry\_at)
    
*   blocked\_user: (expires\_at), (scope); block\_appeal (status)
    
*   suspension: (user\_id, is\_active), (end\_date); suspension\_history (action\_at DESC)
    
*   ban: (expires\_at); ban\_history (action\_at DESC); ban\_evasion (detected\_at)
    
*   org: (parent\_org\_id), (billing\_profile\_id); org\_member (role); org\_seat (status)
    
*   security: device (user\_id, revoked), session (user\_id, revoked), (expires\_at); login\_alert (sent\_at)
    
*   compliance: tax\_profile (validated), residency (review\_status), breach\_notification (sent\_at)
    
*   risk: risk\_signal (user\_id, occurred\_at DESC), risk\_score (score DESC), risk\_hold (until), risk\_alert (sent\_at)
    
*   depth: profile\_depth (depth\_score DESC), profile\_rate\_history (effective\_at DESC)
    
*   audit: (entity\_type, entity\_id, created\_at DESC) + GIN on changes; index actor\_id, correlation\_id
    
*   ops: inbox\_messages (retry\_count DESC)
    

**Partial uniques & checks**

*   Single is\_primary phone/address per user (partial unique).
    
*   language (user\_id, code) unique; skill\_pricing\_tier (skill\_id, tier) unique; profile\_locale\_variant (profile\_id, locale) unique; profile\_rate\_history (profile\_id, effective\_at) unique; profile\_availability\_slot (profile\_id, weekday, start\_minutes, end\_minutes, tz) unique; composite PKs on user\_skill, org\_member, talent\_pool\_member, collection\_item.
    
*   CHECKs: DOB range; experience date range; education year; blocked\_user.blocker\_id <> blocked\_user.blocked\_id; user\_consent.revoked\_at >= accepted\_at when both set.
    
*   Partitioning as noted.
    

**PII lifecycle**

*   Soft deletes via deleted\_at; holds via user\_legal\_hold; exports via data\_export\_job; audit every write.