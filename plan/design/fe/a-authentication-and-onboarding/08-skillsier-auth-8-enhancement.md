## AUTH-8 — Social Login & SSO (Google, Apple, GitHub, Microsoft) + Account Linking / Unlinking + One‑Tap + Enterprise SSO (OIDC/SAML)

> Enhancement add‑on for the existing AUTH‑8 journey in `skillsier-frontend-journeys-claude-final.md`.
> This deliverable fully completes the three sections per the **Enhancement Prompt Template**: **Frontend Implementation**, **Visual Representations**, and **User Stories** (web + mobile), aligned with `combined-fe-folder-structure.md` conventions. Scope covers: *consumer SSO (Google, Apple, GitHub, Microsoft)*, *One‑Tap/Sign‑in with Google*, *account linking/unlinking in Settings*, *provider conflict resolution & merge*, *enterprise SSO discovery by domain (OIDC/SAML)*, *MFA after SSO when required*, *deep‑link mobile handoff*, *analytics*, *a11y*, and *security (PKCE, state, nonce, redirect allowlist)*.

---

### Frontend Implementation

#### A) Page Components (Routes)

**Web (Next.js App Router):**
```
Path: apps/web/app/(auth)/sso/providers/page.tsx
Purpose: Provider selection (Google, Apple, GitHub, Microsoft, Enterprise SSO).
Key Features:
  - One‑tap (optional)
  - Continue with {Provider} buttons
  - “Use email/password” fallback link

Path: apps/web/app/(auth)/sso/callback/page.tsx
Purpose: OAuth/OIDC callback landing; validates PKCE/state/nonce and finalizes login.
Props: search params (code, state, provider)

Path: apps/web/app/(auth)/sso/link/page.tsx
Purpose: Link external identity to an existing account (after login or during conflict resolution).

Path: apps/web/app/(auth)/sso/conflict/page.tsx
Purpose: Conflict UI when provider email matches an existing account; choose link vs new.

Path: apps/web/app/(auth)/sso/enterprise/page.tsx
Purpose: Enterprise SSO discovery by email domain; initiate OIDC/SAML.

Path: apps/web/app/(settings)/(account)/connected-accounts/page.tsx
Purpose: “Connected Accounts” manager (link/unlink, set default).

Path: apps/web/app/(auth)/layout.tsx
Purpose: Shared auth layout (Toaster, ErrorBoundary + Sentry).

Path: apps/web/app/(auth)/sso/*/loading.tsx
Path: apps/web/app/(auth)/sso/*/error.tsx
Purpose: Skeletons and friendly fallbacks.
```

**Mobile (Expo Router):**
```
Path: apps/mobile/app/(public)/(auth)/sso-providers.tsx
Purpose: Provider selection screen.

Path: apps/mobile/app/(public)/(auth)/sso-callback.tsx
Purpose: Mobile OAuth callback (deep link); finalize login.

Path: apps/mobile/app/(public)/(auth)/sso-enterprise.tsx
Purpose: Enterprise discovery (enter email domain).

Path: apps/mobile/app/(tabs)/(authenticated)/(settings)/account/connected-accounts.tsx
Purpose: Manage linked providers after login.
```

> Handoff: AUTH‑8 integrates with **AUTH‑2/AUTH‑3** (sign‑in), **AUTH‑4** (MFA may be required post‑SSO), **AUTH‑6** (step‑up before unlink), and **AUTH‑5/7** (email/phone verification/password reset fallbacks).

---

#### B) Shared Components

**UI Components** (`packages/ui/src/components/`):
```
Component: ProviderButton (Google/Apple/GitHub/Microsoft variants)
Path: packages/ui/src/components/auth/ProviderButton.tsx
Purpose: Consistent provider CTA with brand‑safe visual

Component: ProviderList
Path: packages/ui/src/components/auth/ProviderList.tsx
Purpose: Renders buttons + optional One‑Tap area

Component: OneTapGoogle
Path: packages/ui/src/components/auth/OneTapGoogle.tsx
Purpose: Wraps Google One‑Tap; emits credential response

Component: ConflictCard
Path: packages/ui/src/components/auth/ConflictCard.tsx
Purpose: Explain email collision; options to link or continue separate

Component: ConnectedAccountCard
Path: packages/ui/src/components/settings/ConnectedAccountCard.tsx
Purpose: Shows linked provider, last used, make default, unlink

Component: DangerDialog
Path: packages/ui/src/components/common/DangerDialog.tsx
Purpose: Confirm unlink; explains fallback requirements
```

**Feature Components** (`packages/lib/src/features/sso/components/`):
```
Component: SsoStart.tsx
Purpose: Create PKCE, state, nonce; redirect to provider auth URL

Component: SsoCallbackFinalizer.tsx
Purpose: Exchange code for tokens; validate state/nonce; sign in

Component: SsoEnterpriseDiscovery.tsx
Purpose: Domain discovery → OIDC/SAML redirect builder

Component: SsoLinkForm.tsx
Purpose: Link provider to current user (step‑up if needed)

Component: ConnectedAccountsPanel.tsx
Purpose: Listing + actions for linked identities
```

**Domain‑Specific Components:**
```
Component: SsoAftercareNudge.tsx
Purpose: Nudges to enable MFA/passkey after SSO sign‑in
```

---

#### C) Hooks & Data Fetching

**TanStack Query / Mutations** (`packages/lib/src/features/sso/hooks/`):
```
Hook: useSsoStart.ts
POST /v1/auth/sso/start { provider, intent?, redirectUri }
→ { authUrl, pkce:{ codeChallenge, method }, state, nonce }

Hook: useSsoCallback.ts
POST /v1/auth/sso/callback { provider, code, codeVerifier, state, nonce }
→ { user, session, requiresMfa?:boolean, isNew?:boolean }

Hook: useSsoLink.ts
POST /v1/auth/sso/link { provider, code?, idToken?, codeVerifier?, state?, nonce? }
→ { linked:true }

Hook: useSsoUnlink.ts
POST /v1/auth/sso/unlink { providerId }
→ { unlinked:true }

Hook: useConnectedAccountsQuery.ts
GET /v1/auth/sso/accounts
→ { accounts:[{ id, provider, emailMasked, lastUsedAt, isDefault }] }

Hook: useSsoEnterpriseDiscovery.ts
POST /v1/auth/sso/enterprise/discover { emailOrDomain }
→ { provider:"oidc"|"saml", idpName, authUrl }

Hook: useOneTapInit.ts
POST /v1/auth/sso/google/onetap/init
→ { clientId, allowedOrigins }
```

**State Management (Zustand)** (`packages/lib/src/stores/sso-store.ts`):
```
interface SsoState {
  pkce?: { codeVerifier:string, codeChallenge:string, method:"S256"|"plain" };
  state?: string;
  nonce?: string;
  intent?: string|null;     // e.g., "signin" | "link"
  postAuthRedirect?: string|null;
  set: (p: Partial<SsoState>) => void;
  clear: () => void;
}
```

---

#### D) Utilities & Helpers

**Validation Schemas (Zod)** (`packages/lib/src/schemas/sso/`):
```
ssoStartSchema.ts        → provider enum, redirectUri, intent
ssoCallbackSchema.ts     → provider, code, state, nonce, codeVerifier
enterpriseDiscover.ts    → email/domain validation
```

**Security & OAuth Utils** (`packages/lib/src/utils/`):
```
oauth/pkce.ts         → create code_verifier + code_challenge (S256)
oauth/state.ts        → random state; session binding
oauth/nonce.ts        → OIDC nonce generation
oauth/redirect.ts     → allowlist + canonical redirect resolver
link/deeplink.ts      → app deep‑link helpers (scheme, universal links)
```

**Types** (`packages/lib/src/types/sso.ts`):
```
Provider = "google"|"apple"|"github"|"microsoft"|"enterprise"
SsoStartResponse, SsoCallbackResponse, ConnectedAccount, EnterpriseDiscovery
```

---

#### E) Mobile‑Specific Components

**Navigation Components:**
```
Component: SsoStack (Expo Router)
Paths:
  - /(public)/(auth)/sso-providers
  - /(public)/(auth)/sso-callback
  - /(public)/(auth)/sso-enterprise
  - /(tabs)/(authenticated)/(settings)/account/connected-accounts
```

**Native Features:**
```
- ASWebAuthenticationSession (iOS) / Custom Tabs (Android) for OAuth
- expo-linking for deep link callback
- SecureStore for ephemeral PKCE/state/nonce (not persisted across reinstalls)
```

---

#### F) Layout Components

**Web:**
```
Layout: (auth) Layout
Path: apps/web/app/(auth)/layout.tsx
Features: Minimal chrome, Toaster, ErrorBoundary + Sentry
```

**Mobile:**
```
Layout: (public)/(auth)/_layout.tsx
Features: Header/back; modal for conflict/help
```

---

#### G) Error Boundaries & Loading States

**Error Handling:**
```
Routes: (auth)/sso/*/error.tsx → friendly fallbacks (state/nonce mismatch, denied consent)
Components: InlineError banners in conflict/link pages
```

**Loading States:**
```
Routes: (auth)/sso/*/loading.tsx → skeletons/spinners
Components: Progress bar on callback exchange
```

---

### Visual Representations

#### Screen 1 — Continue with a Provider (Selection)

**Web View**
```
┌─────────────────────────────────────────────────────────────────────────────┐
│ Sign in to Skillsier                                                        │
├─────────────────────────────────────────────────────────────────────────────┤
│ Continue with                                                               │
│ ┌───────────────────────────────────────────────────────────────────────┐   │
│ │  ⓖ  Continue with Google                                              │   │
│ └───────────────────────────────────────────────────────────────────────┘   │
│ ┌───────────────────────────────────────────────────────────────────────┐   │
│ │    Continue with Apple                                               │   │
│ └───────────────────────────────────────────────────────────────────────┘   │
│ ┌───────────────────────────────────────────────────────────────────────┐   │
│ │    Continue with GitHub                                              │   │
│ └───────────────────────────────────────────────────────────────────────┘   │
│ ┌───────────────────────────────────────────────────────────────────────┐   │
│ │  Ⓜ  Continue with Microsoft                                           │   │
│ └───────────────────────────────────────────────────────────────────────┘   │
│                                                                            │
│ Enterprise SSO?  Enter your work email → [ ________ ] [ Continue ]         │
│                                                                            │
│ Prefer password?  [ Use email & password ]                                 │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Mobile View**
```
┌──────────────────────────┐
│ ← Sign in                │
├──────────────────────────┤
│ ⓖ  Continue with Google  │
│   Continue with Apple   │
│   Continue with GitHub  │
│ Ⓜ  Continue with MSFT    │
│                          │
│ Enterprise SSO           │
│ [ you@company.com  ] →   │
│                          │
│ Use email & password     │
└──────────────────────────┘
```

---

#### Screen 2 — OAuth Callback (Finalizing)

```
┌──────────────────────────────────────────────┐
│ Completing sign‑in…                          │
├──────────────────────────────────────────────┤
│ • Verifying state & nonce                    │
│ • Exchanging code (PKCE)                     │
│ • Creating session                           │
│ ✅ Success!                                   │
│ [ Continue ]                                  │
└──────────────────────────────────────────────┘
```

**Error (State/Nonce Mismatch)**
```
┌──────────────────────────────────────────────┐
│ ❌ We couldn’t verify your sign‑in.          │
│ Reason: Security check failed.               │
│ [ Try again ]  [ Use email & password ]      │
└──────────────────────────────────────────────┘
```

---

#### Screen 3 — Conflict Resolution (Email Already In Use)

```
┌──────────────────────────────────────────────┐
│ This email is already used on Skillsier      │
├──────────────────────────────────────────────┤
│ You can link your Google account to          │
│ your existing Skillsier account.             │
│                                              │
│ [ Link to my account ]   [ Create separate ] │
│ [ Use email & password ] (I’ll link later)   │
└──────────────────────────────────────────────┘
```

---

#### Screen 4 — Connected Accounts (Settings)

```
┌──────────────────────────────────────────────┐
│ Connected Accounts                           │
├──────────────────────────────────────────────┤
│ Google      you@exa…   Default ✓   [ Unlink ] [ Make default ]            │
│ Apple       —          Not linked  [ Link ]                                │
│ GitHub      dev@exa…   Last used: 2d  [ Unlink ]                           │
│ Microsoft   —          Not linked  [ Link ]                                │
│                                                                          │
│ Tip: Enable MFA to add extra protection.                                  │
└──────────────────────────────────────────────┘
```

---

#### Screen 5 — Enterprise SSO Discovery

```
┌──────────────────────────────────────────────┐
│ Enterprise SSO                               │
├──────────────────────────────────────────────┤
│ Work email  [ you@company.com ]              │
│ [ Continue ]                                  │
│ This will redirect you to your company’s IdP. │
└──────────────────────────────────────────────┘
```

---

### User Stories — Complete

> Abbrev: **W**=Web, **M**=Mobile. Each story follows “As a … I want … so that …” and includes components, routes, APIs, with **≥5 Happy** & **≥10 Bad** scenarios.

### Story 1 — Select a Social Provider

**Story**
```
As a New or Returning User
I want to pick a social or enterprise SSO provider
So that I can sign in without typing a password
```

**Components & Routes**
- **W:** `(auth)/sso/providers/page.tsx` → `ProviderList`, `ProviderButton`, `OneTapGoogle`, `SsoStart`
- **M:** `/(public)/(auth)/sso-providers.tsx`
- **APIs:** `POST /v1/auth/sso/start`

**Acceptance Criteria — ✅ Happy**
1. **AC1.1** Provider buttons initiate PKCE/state/nonce and redirect to the provider.
2. **AC1.2** One‑Tap shows where supported and user‑enabled.
3. **AC1.3** Enterprise email field validates domain format.
4. **AC1.4** “Use email & password” link navigates to AUTH‑2.
5. **AC1.5** Analytics `sso.start` with provider captured (no PII).

**Acceptance Criteria — ❌ Bad**
1. **AC1.6** Unsupported browser blocks One‑Tap → silently hides widget.
2. **AC1.7** Rate‑limit on start → toast + disabled button briefly.
3. **AC1.8** PKCE generation fails → retry with new verifier.
4. **AC1.9** Blocked third‑party cookies (Apple) → guidance fallback.
5. **AC1.10** Redirect URI not allow‑listed → hard error and log.
6. **AC1.11** Network/5xx → retry and keep selection intact.
7. **AC1.12** Localization missing → English fallback.
8. **AC1.13** A11y: buttons labeled with provider names, focus visible.
9. **AC1.14** Screen too small → buttons stack vertically responsively.
10. **AC1.15** Analytics disabled → non‑blocking; queue drops.


---

### Story 2 — Complete OAuth/OIDC Callback (Web)

**Story**
```
As a User arriving from the provider
I want Skillsier to finalize my sign‑in securely
So that I can access my account
```

**Components & Routes**
- **W:** `(auth)/sso/callback/page.tsx` → `SsoCallbackFinalizer`
- **APIs:** `POST /v1/auth/sso/callback`

**Acceptance Criteria — ✅ Happy**
1. **AC2.1** State & nonce validated; PKCE exchange succeeds.
2. **AC2.2** New session created; redirect to last intent or dashboard.
3. **AC2.3** If `requiresMfa`, route to AUTH‑2 challenge step.
4. **AC2.4** Analytics `sso.callback.success` logged once.
5. **AC2.5** Tokens not stored in localStorage; httpOnly cookies used.

**Acceptance Criteria — ❌ Bad**
1. **AC2.6** State mismatch → fail with clear retry path.
2. **AC2.7** Nonce mismatch → hard error; suggest retry.
3. **AC2.8** Code already used → idempotent success or error with restart.
4. **AC2.9** Provider error=access_denied → user can pick another provider.
5. **AC2.10** Clock skew errors → tolerant window; guidance text.
6. **AC2.11** Network/5xx → retry; preserve intent.
7. **AC2.12** CSRF issues → refresh then retry.
8. **AC2.13** Token audience/scope invalid → show support path.
9. **AC2.14** Provider down → fallback to email/password.
10. **AC2.15** PII leakage in logs blocked by scrubbers.


---

### Story 3 — Mobile OAuth Flow & Deep Link

**Story**
```
As a Mobile User
I want to authenticate with a provider through the system browser
So that I can sign in securely with deep‑link handoff
```

**Components & Routes**
- **M:** `sso-providers.tsx` (start), `sso-callback.tsx` (deep link)
- **APIs:** same as web

**Acceptance Criteria — ✅ Happy**
1. **AC3.1** Uses ASWebAuthenticationSession/Custom Tabs; sets PKCE/state/nonce.
2. **AC3.2** Deep link returns to app and finalizes sign‑in.
3. **AC3.3** If app not installed → mobile web fallback works.
4. **AC3.4** Analytics `sso.mobile.success` logged.
5. **AC3.5** SecureStore used for ephemeral PKCE/state/nonce. 

**Acceptance Criteria — ❌ Bad**
1. **AC3.6** Deep link misconfigured → shows QR to open in app, or web fallback.
2. **AC3.7** User cancels browser → return to selection.
3. **AC3.8** Intent chooser confusion → explain “Open with Skillsier”.
4. **AC3.9** Old app cannot handle callback → open in web.
5. **AC3.10** Network/5xx → retry; keep intent.
6. **AC3.11** State lost (OS reclaimed) → restart cleanly.
7. **AC3.12** Token left in URL → scrub via POST‑redirect, never log.
8. **AC3.13** Insecure scheme → only https/universal links allowed.
9. **AC3.14** App killed before callback → resume and prompt to retry.
10. **AC3.15** Analytics blocked → non‑blocking.


---

### Story 4 — Email Collision & Account Linking (During Sign‑in)

**Story**
```
As a Returning User whose email matches an existing account
I want to link my provider identity to my account
So that I don’t create duplicates
```

**Components & Routes**
- **W:** `(auth)/sso/conflict/page.tsx` → `ConflictCard`, `SsoLinkForm`
- **M:** handled within callback & a modal card
- **APIs:** `POST /v1/auth/sso/link`

**Acceptance Criteria — ✅ Happy**
1. **AC4.1** Conflict explains options: link to existing or create separate.
2. **AC4.2** Choosing “Link” requires step‑up (AUTH‑6) or email verify (AUTH‑5).
3. **AC4.3** On success, identities merged; user lands on dashboard.
4. **AC4.4** Analytics `sso.link.success` logged.
5. **AC4.5** Future sign‑ins with this provider go directly to account.

**Acceptance Criteria — ❌ Bad**
1. **AC4.6** Linking denied by org policy → copy explains; offer email login.
2. **AC4.7** Different email on provider vs account → require explicit confirmation.
3. **AC4.8** Re‑auth required and fails → remain on conflict with helper text.
4. **AC4.9** Attempt to link to another user’s account → blocked and audited.
5. **AC4.10** Provider not returning verified email → ask user to confirm.
6. **AC4.11** Network/5xx → retry with backoff.
7. **AC4.12** Rate‑limit linking → cool‑down messaging.
8. **AC4.13** Duplicate clicks → idempotent action.
9. **AC4.14** A11y: buttons labeled; focus order sensible.
10. **AC4.15** i18n keys missing → fallback strings.


---

### Story 5 — Link a Provider in Settings (Post‑Login)

**Story**
```
As a Signed‑in User
I want to connect a provider to my account
So that I can sign in more easily next time
```

**Components & Routes**
- **W:** `(settings)/(account)/connected-accounts/page.tsx` → `ConnectedAccountsPanel`, `ProviderButton`
- **M:** `…/connected-accounts.tsx`
- **APIs:** `POST /v1/auth/sso/start` (intent=link), then `POST /v1/auth/sso/link`

**Acceptance Criteria — ✅ Happy**
1. **AC5.1** Click “Link Google” → OAuth → returns linked in list.
2. **AC5.2** Step‑up (AUTH‑6) may be required; then proceed.
3. **AC5.3** Mark one provider as default sign‑in method.
4. **AC5.4** Success toast; last‑used timestamp updates.
5. **AC5.5** Analytics `sso.settings.link.success` logged.

**Acceptance Criteria — ❌ Bad**
1. **AC5.6** Already linked → button disabled with note.
2. **AC5.7** Provider returns different email → confirmation required.
3. **AC5.8** Network/5xx → retry; UI state preserved.
4. **AC5.9** Policy forbids multiple providers → explain and enforce.
5. **AC5.10** A11y: cards reachable; buttons have aria‑labels.
6. **AC5.11** i18n fallback on provider names/translations.
7. **AC5.12** Rate‑limit → cool‑down with timer.
8. **AC5.13** Revoked consent at provider → handle and show fix steps.
9. **AC5.14** IdP returns no email → request user email verification (AUTH‑5).
10. **AC5.15** Analytics blocked → non‑blocking.


---

### Story 6 — Unlink a Provider (With Fallback Credential Guard)

**Story**
```
As a Security‑conscious User
I want to disconnect a linked provider safely
So that I don’t lose access to my account
```

**Components & Routes**
- **W/M:** `ConnectedAccountCard` → `DangerDialog`, `useSsoUnlink`
- **APIs:** `POST /v1/auth/sso/unlink`

**Acceptance Criteria — ✅ Happy**
1. **AC6.1** If no other login method exists, UI requires setting a password or passkey first.
2. **AC6.2** Step‑up required before unlink; success updates list.
3. **AC6.3** Email confirmation of unlink sent.
4. **AC6.4** Analytics `sso.unlink.success` logged.
5. **AC6.5** Audit entry with masked data created.

**Acceptance Criteria — ❌ Bad**
1. **AC6.6** Only linked method → unlink blocked until fallback set.
2. **AC6.7** Step‑up fails → no unlink; guidance shown.
3. **AC6.8** Network/5xx → retry; state intact.
4. **AC6.9** Race with simultaneous unlink on another device → idempotent.
5. **AC6.10** Provider still shows as linked due to cache → refetch.
6. **AC6.11** A11y: dialog focus trapped; ESC closes; labels present.
7. **AC6.12** i18n fallback strings used if missing.
8. **AC6.13** Policy disallows unlink (enterprise) → explain.
9. **AC6.14** Session expired → re‑auth then continue.
10. **AC6.15** Analytics blocked → non‑blocking.


---

### Story 7 — One‑Tap / Sign in with Google

**Story**
```
As a User on supported browsers
I want One‑Tap sign‑in without page navigation
So that I can authenticate quickly
```

**Components & Routes**
- **W:** `OneTapGoogle` + `useOneTapInit`
- **APIs:** `POST /v1/auth/sso/google/onetap/init`, `POST /v1/auth/sso/callback` with id_token

**Acceptance Criteria — ✅ Happy**
1. **AC7.1** One‑Tap shows only when allowed and not dismissed recently.
2. **AC7.2** Successful credential posts id_token to callback; session created.
3. **AC7.3** Handles guest → new account or link if collision.
4. **AC7.4** Opt‑out remembered for a period.
5. **AC7.5** Analytics `sso.onetap.success` logged.

**Acceptance Criteria — ❌ Bad**
1. **AC7.6** Third‑party cookies disabled → One‑Tap hidden.
2. **AC7.7** id_token invalid → error; suggest standard redirect.
3. **AC7.8** Collision requires link → route to conflict screen.
4. **AC7.9** A11y: escape and screen reader announce.
5. **AC7.10** Rate‑limit events; avoid spam after dismissal.
6. **AC7.11** Network/5xx → retry gracefully.
7. **AC7.12** Malicious frame → CSP blocks; no X‑frame allowed.
8. **AC7.13** PII in logs → scrubbed.
9. **AC7.14** Consent refused → user continues without One‑Tap.
10. **AC7.15** Analytics blocked → non‑blocking.


---

### Story 8 — Enterprise SSO Discovery & Redirect

**Story**
```
As an Enterprise User
I want to be routed to my company’s IdP based on my email
So that I can authenticate with corporate credentials
```

**Components & Routes**
- **W:** `(auth)/sso/enterprise/page.tsx` → `SsoEnterpriseDiscovery`
- **M:** `/(public)/(auth)/sso-enterprise.tsx`
- **APIs:** `POST /v1/auth/sso/enterprise/discover`

**Acceptance Criteria — ✅ Happy**
1. **AC8.1** Entering email shows IdP name/logo and redirects to OIDC/SAML authUrl.
2. **AC8.2** Supports direct domain input; validates MX/known config (optional).
3. **AC8.3** Post‑auth returns to callback like normal providers.
4. **AC8.4** Analytics `sso.enterprise.start` captured.
5. **AC8.5** Copy clarifies data sharing and consent.

**Acceptance Criteria — ❌ Bad**
1. **AC8.6** Unknown domain → guide to admin contact or use standard providers.
2. **AC8.7** IdP metadata mismatch → safe error and support path.
3. **AC8.8** SAML clock skew → tolerant window; clear message.
4. **AC8.9** OIDC audience mismatch → error and log.
5. **AC8.10** Network/5xx → retry; keep entered domain.
6. **AC8.11** Policy forbids self‑serve enterprise → hide or request invite.
7. **AC8.12** i18n missing → fallback strings.
8. **AC8.13** A11y: labels and help text accessible.
9. **AC8.14** Redirect URI mismatch → blocked; log incident.
10. **AC8.15** Analytics blocked → non‑blocking.


---

### Story 9 — MFA After SSO (Conditional)

**Story**
```
As a Security‑hardened Workspace
I want SSO sign‑ins to still require MFA when policy demands
So that risk is reduced
```

**Components & Routes**
- **W/M:** Hand off to AUTH‑2 challenge step after callback when `requiresMfa`
- **APIs:** same as AUTH‑2 hooks

**Acceptance Criteria — ✅ Happy**
1. **AC9.1** If policy requires, users complete MFA after SSO before access.
2. **AC9.2** MFA enrollment prompt (AUTH‑4) shown if no factor exists and policy allows grace.
3. **AC9.3** Remember‑device honored for the MFA window.
4. **AC9.4** Analytics `sso.mfa.required` logged.
5. **AC9.5** Back navigation does not bypass MFA gate.

**Acceptance Criteria — ❌ Bad**
1. **AC9.6** MFA service down → fallback guidance and support path.
2. **AC9.7** Factor removed mid‑flow → choose alternate factor.
3. **AC9.8** Rate‑limit on codes → timers and messages.
4. **AC9.9** App reload → resume MFA step with preserved intent.
5. **AC9.10** i18n missing → fallback.
6. **AC9.11** Analytics blocked → non‑blocking.
7. **AC9.12** A11y: code inputs announced correctly.
8. **AC9.13** Offline → wait to reconnect; safe banner.
9. **AC9.14** Time drift → rely on server time.
10. **AC9.15** Multiple tabs show different steps → single source of truth.


---

### Story 10 — Set Default Sign‑In Method

**Story**
```
As a User with multiple linked providers
I want to choose my default sign‑in option
So that sign‑in is faster next time
```

**Components & Routes**
- **W/M:** `(settings)/(account)/connected-accounts` → `ConnectedAccountsPanel`
- **APIs:** `POST /v1/auth/sso/accounts/default`

**Acceptance Criteria — ✅ Happy**
1. **AC10.1** Selecting “Make default” updates server and UI.
2. **AC10.2** Start screen jumps straight to default on next visit (optional).
3. **AC10.3** Email notification confirms change (optional).
4. **AC10.4** Analytics `sso.default.changed` recorded.
5. **AC10.5** Works consistently across devices after sync.

**Acceptance Criteria — ❌ Bad**
1. **AC10.6** Server rejects default → message and revert UI.
2. **AC10.7** Policy forbids certain defaults → explain and enforce.
3. **AC10.8** Offline → queued and retried; UI shows pending.
4. **AC10.9** Race with unlink → serialize; resolve conflict.
5. **AC10.10** i18n fallback if missing.
6. **AC10.11** A11y: radio group or buttons accessible.
7. **AC10.12** Analytics blocked → non‑blocking.
8. **AC10.13** Cache stale after change → refetch on focus.
9. **AC10.14** Multiple tabs set different defaults → last write wins.
10. **AC10.15** Provider revoked at IdP → default auto‑cleared.


---

### Story 11 — Analytics & Telemetry

**Story**
```
As a Product Team
I want to measure SSO funnel performance across providers
So that we can improve conversion
```

**Events (examples)**
- `sso.start` / `.callback.success` / `.callback.error`
- `sso.link.success` / `.unlink.success`
- `sso.mobile.success`
- `sso.enterprise.start` / `.success`
- `sso.default.changed`

**Acceptance Criteria — ✅ Happy**
1. **AC11.1** Events debounced/idempotent; no duplicates on rerender.
2. **AC11.2** No PII (tokens/emails) in payloads.
3. **AC11.3** Client queues and retries analytics on failures.
4. **AC11.4** Versioned schemas validated in CI.
5. **AC11.5** Provider breakdown dashboards by device/platform.

**Acceptance Criteria — ❌ Bad**
1. **AC11.6** Payload includes tokens → blocked by linter.
2. **AC11.7** Missing success on deep link → app‑open hook ensures send.
3. **AC11.8** Time skew → server timestamp included.
4. **AC11.9** Ad‑blockers drop beacons → non‑blocking; drop safe.
5. **AC11.10** Event names change without version → build fails.
6. **AC11.11** Over‑instrumentation → perf budget enforced.
7. **AC11.12** Analytics library crashes → sandbox & try/catch.
8. **AC11.13** Duplicates from multiple tabs → storage‑based dedupe.
9. **AC11.14** PII logged via console → disabled in prod.
10. **AC11.15** GDPR: unconsented tracking → respect consent gates.


---

### Story 12 — Security Hardening (PKCE, State, Nonce, Redirect Allowlist)

**Story**
```
As a Security Team
I want robust OAuth/OIDC protections
So that SSO flows resist common attacks
```

**Acceptance Criteria — ✅ Happy**
1. **AC12.1** PKCE S256 used for public clients; verifier stored ephemerally.
2. **AC12.2** State and nonce bound to session; validated on callback.
3. **AC12.3** Redirect URIs are allow‑listed and canonicalized.
4. **AC12.4** httpOnly, sameSite cookies used; no tokens in localStorage.
5. **AC12.5** Provider public keys (JWKS) cached & rotated safely.

**Acceptance Criteria — ❌ Bad**
1. **AC12.6** State/nonce mismatch → blocked with clear guidance.
2. **AC12.7** Code interception → PKCE prevents; audit ensured.
3. **AC12.8** Open redirect via next param → allowlist enforced.
4. **AC12.9** Log tokens or id_tokens → scrub & deny.
5. **AC12.10** Apple nonce omitted → callback rejected.
6. **AC12.11** GitHub scope too broad → reduced to least privilege.
7. **AC12.12** Clock skew breaks validation → tolerant checks.
8. **AC12.13** JWKS fetching fails → retry; safe cache fallback.
9. **AC12.14** Third‑party script modifies oauth params → CSP prevents.
10. **AC12.15** SameSite issues on Safari → compatible cookie strategy.


---

## Non‑Functional Requirements (AUTH‑8 scope)

- **Performance:** SSO start < 400ms p95; callback finalize < 600ms p95
- **Security:** PKCE S256; state/nonce; redirect allowlist; httpOnly cookies; CSP; audit logging; Sentry
- **Privacy:** No tokens/PII in logs or analytics; masked emails in UI
- **Compatibility:** Web (Chrome/Firefox/Safari/Edge), iOS 14+, Android 10+; provider‑specific quirks handled
- **Testing:** Unit tests (schemas/hooks/utils), e2e per provider (web + mobile), deep‑link tests, a11y (axe)

---

**End of AUTH‑8 Enhancements (Complete).**
