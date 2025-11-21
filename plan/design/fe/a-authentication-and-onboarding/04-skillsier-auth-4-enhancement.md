## AUTH-4 — MFA Enrollment & Account Security (TOTP / SMS / Email, Passkeys, Recovery Codes, Trusted Devices, Policy Enforcement)

> Enhancement add‑on for the existing AUTH‑4 journey in `skillsier-frontend-journeys-claude-final.md`.  
> This deliverable fully completes the three required sections: **Frontend Implementation**, **Visual Representations**, and **User Stories** (web + mobile), aligned with `combined-fe-folder-strucure.md` path conventions. Scope: enable/disable MFA methods, TOTP QR pairing, SMS/email codes, passkeys (WebAuthn/FIDO2), recovery codes management, trusted devices, and MFA policy enforcement.

---

### Frontend Implementation

#### A) Page Components (Routes)

**Web (Next.js App Router):**
```
Path: apps/web/app/(settings)/(security)/mfa/page.tsx
Purpose: Security overview & MFA hub (status, methods, recovery codes, devices).
Props: None (client component with data fetching).
Key Features:
  - Lists enabled/available MFA methods
  - Entry points to setup wizards
  - Recovery codes generate/view
  - Trusted devices list & revoke
  - Org policy notice (if enforced)

Path: apps/web/app/(settings)/(security)/mfa/setup/totp/page.tsx
Purpose: TOTP enrollment wizard (QR + secret → verify code → success).
Props: None
Key Features:
  - Shows issuer/account, QR, plain secret
  - Code input with auto‑advance
  - Backup instructions

Path: apps/web/app/(settings)/(security)/mfa/setup/sms/page.tsx
Purpose: SMS enrollment wizard (phone → verify → success).
Props: None
Key Features:
  - E.164 phone input with country picker
  - Resend with cooldown
  - Trust device option (post‑enroll)

Path: apps/web/app/(settings)/(security)/mfa/setup/passkey/page.tsx
Purpose: WebAuthn/FIDO2 security key or platform passkey enrollment.
Props: None
Key Features:
  - Start/finish WebAuthn registration
  - Nickname credential
  - Success + test sign‑in

Path: apps/web/app/(settings)/(security)/mfa/recovery-codes/page.tsx
Purpose: Generate/display/rotate recovery codes; download/print.
Props: None
Key Features:
  - One‑time reveal with confirm
  - Rotate invalidates old set
  - Print/download (PDF/CSV)

Path: apps/web/app/(settings)/(security)/mfa/disable/[method]/page.tsx
Purpose: Disable/remove an MFA method (danger action with confirm).
Props: method (totp|sms|passkey|email)
Key Features:
  - Confirm modal + re‑auth prompt
  - Requires an available second factor

Path: apps/web/app/(settings)/(security)/mfa/devices/page.tsx
Purpose: Trusted devices (remembered) management and revocation.
Props: None
Key Features:
  - Device list with last seen/IP/UA
  - Revoke single/all

Path: apps/web/app/(settings)/(security)/mfa/policy/page.tsx
Purpose: Policy interstitial if org enforces MFA; guides to complete setup.
Props: redirect (where to go after compliance)
```

**Mobile (Expo Router):**
```
Path: apps/mobile/app/(tabs)/(authenticated)/(settings)/security/mfa.tsx
Purpose: Security overview & MFA hub.
Key Features:
  - Same capabilities as web, mobile‑optimized

Path: apps/mobile/app/(tabs)/(authenticated)/(settings)/security/mfa-setup-totp.tsx
Path: apps/mobile/app/(tabs)/(authenticated)/(settings)/security/mfa-setup-sms.tsx
Path: apps/mobile/app/(tabs)/(authenticated)/(settings)/security/mfa-setup-passkey.tsx
Path: apps/mobile/app/(tabs)/(authenticated)/(settings)/security/mfa-recovery-codes.tsx
Path: apps/mobile/app/(tabs)/(authenticated)/(settings)/security/mfa-devices.tsx
Path: apps/mobile/app/(tabs)/(authenticated)/(settings)/security/mfa-policy.tsx
```

> Handoff: AUTH‑4 is post‑login security management. It integrates with **AUTH‑2** (MFA challenge at sign‑in) and **ACC/Settings** modules.

---

#### B) Shared Components

**UI Components** (`packages/ui/src/components/`):
```
Component: Button / LinkButton
Purpose: Primary/secondary CTAs (Set up, Verify, Disable, Revoke)
Props: variant, size, isLoading, leftIcon, rightIcon

Component: Input
Purpose: Text fields (labels, helper, error)
Props: id, type, value, onChange, error, autoComplete

Component: PhoneInput
Path: packages/ui/src/components/forms/PhoneInput.tsx
Purpose: Country picker + E.164 formatting

Component: CodeInput
Path: packages/ui/src/components/forms/CodeInput.tsx
Purpose: 6‑digit code entry with auto‑advance and paste support

Component: QRCode
Path: packages/ui/src/components/media/QRCode.tsx
Purpose: Render TOTP QR (otpauth://)

Component: Toggle / Switch
Purpose: Enable/disable options (remember device)

Component: Card, List, Badge, Tooltip, Modal, Drawer, Alert/Banner, Table, EmptyState, Spinner
```

**Feature Components** (`packages/lib/src/features/security/components/`):
```
Component: MfaOverview.tsx
Purpose: Fetch & render methods, recovery status, trusted devices count

Component: TotpSetupWizard.tsx
Purpose: Step 1 (QR/secret) → Step 2 (code verify) → Step 3 (success)

Component: SmsSetupWizard.tsx
Purpose: Phone entry → code verify → success (resend, cooldown)

Component: PasskeyEnrollPanel.tsx
Purpose: Start/finish WebAuthn registration; nickname field

Component: RecoveryCodesPanel.tsx
Purpose: Generate, reveal, rotate, download/print codes

Component: TrustedDevicesTable.tsx
Purpose: List + Revoke device(s); confirm dialogs

Component: DisableMethodDialog.tsx
Purpose: Danger confirm; re‑auth prompt (password or passkey)

Component: PolicyInterstitial.tsx
Purpose: Enforced MFA policy walkthrough with progress
```

**Domain‑Specific Components**
```
Component: SecurityHero.tsx (web)
Purpose: Informational header and policy status
```

---

#### C) Hooks & Data Fetching

**TanStack Query / Mutations** (`packages/lib/src/features/security/hooks/`):
```
Hook: useMfaMethodsQuery.ts
GET /v1/auth/mfa/methods
→ { enabled: Method[], available: Method[], policy: {...} }

Hook: useTotpEnrollStart.ts
POST /v1/auth/mfa/totp/enroll/start
→ { secret, issuer, account, otpauthUrl }

Hook: useTotpEnrollVerify.ts
POST /v1/auth/mfa/totp/enroll/verify { code }
→ { success: true }

Hook: useSmsEnrollStart.ts
POST /v1/auth/mfa/sms/enroll/start { phone }
→ { challengeId, resendAt }

Hook: useSmsEnrollVerify.ts
POST /v1/auth/mfa/sms/enroll/verify { challengeId, code }
→ { success: true }

Hook: usePasskeyRegisterStart.ts
POST /v1/auth/webauthn/register/start
→ { publicKeyCredentialCreationOptions }

Hook: usePasskeyRegisterFinish.ts
POST /v1/auth/webauthn/register/finish { credential }
→ { success: true, credentialId, nickname }

Hook: useRecoveryCodesGenerate.ts
POST /v1/auth/mfa/recovery/generate
→ { codes: string[] }

Hook: useTrustedDevicesQuery.ts
GET /v1/auth/trusted-devices
→ { devices: { id, ua, ip, lastSeen, createdAt }[] }

Hook: useRevokeDeviceMutation.ts
DELETE /v1/auth/trusted-devices/{id}

Hook: useMfaDisableMethod.ts
POST /v1/auth/mfa/disable { method }

Hook: useMfaPolicyQuery.ts
GET /v1/auth/mfa/policy
→ { required: boolean, dueBy?: string }

Hook: useResendSmsCode.ts
POST /v1/auth/mfa/sms/enroll/resend { challengeId }
```

**State Management (Zustand)** (`packages/lib/src/stores/security-store.ts`):
```
interface SecurityState {
  activeWizard?: "totp"|"sms"|"passkey"|null;
  challengeId?: string|null;
  cooldownEndsAt?: number|null;
  methods?: string[];
  set: (p: Partial<SecurityState>) => void;
  clear: () => void;
}
```

---

#### D) Utilities & Helpers

**Validation Schemas (Zod)** (`packages/lib/src/schemas/security/`):
```
phoneSchema.ts      → E.164 phone validation
codeSchema.ts       → 6 digits numeric
nicknameSchema.ts   → 1–40 chars, printables
```

**Security & Formatting Utils** (`packages/lib/src/utils/`):
```
security/webauthn.ts     → wrap navigator.credentials.create()
security/totp.ts         → otpauth URL builder (client‑side only for previews)
format/maskPhone.ts      → +1 ***‑***‑1234
time/cooldown.ts         → server‑synced countdown helpers
```

**Types** (`packages/lib/src/types/security.ts`):
```
MfaMethod, MfaPolicy, TrustedDevice, RecoveryCodesResponse, PasskeyCredential
```

---

#### E) Mobile‑Specific Components

**Navigation Components:**
```
Component: SettingsSecurityStack (Expo Router)
Path: apps/mobile/app/(tabs)/(authenticated)/(settings)/security/_layout.tsx
Screens:
  - mfa
  - mfa-setup-totp
  - mfa-setup-sms
  - mfa-setup-passkey
  - mfa-recovery-codes
  - mfa-devices
  - mfa-policy
```

**Native Features:**
```
- Push haptic feedback on verify/success
- Secure storage for local MFA hints (never secrets)
- WebAuthn via react-native-passkeys (platform‑dependent)
```

---

#### F) Layout Components

**Web:**
```
Layout: (settings)/(security) Layout
Path: apps/web/app/(settings)/(security)/layout.tsx
Features:
  - Settings sidebar navigation
  - Section header & breadcrumbs
  - ErrorBoundary + Sentry
  - <Toaster/> mount
```

**Mobile:**
```
Layout: (tabs)/(authenticated)/(settings)/security/_layout.tsx
Features:
  - Header/back
  - Modals for confirm dialogs
```

---

#### G) Error Boundaries & Loading States

**Error Handling:**
```
Routes: .../mfa/*/error.tsx
Purpose: Route fallback; reset; Sentry capture

Components: DisableMethodDialog, VerifyErrorBanner
Purpose: Render API errors, rescues, guidance
```

**Loading States:**
```
Routes: .../mfa/*/loading.tsx → skeletons
Component: Spinner.tsx → inline busy states
```

---

### Visual Representations

#### Screen 1 — Security & MFA Overview

**Web View (1280–1440px):**
```
┌─────────────────────────────────────────────────────────────────────────────┐
│ Settings ▸ Security                                                         │
├─────────────────────────────────────────────────────────────────────────────┤
│ Security & MFA                                                              │
│ ─────────────────────────────────────────────────────────────────────────   │
│ MFA Status:  Enabled  [Badge ✓]   Policy:  Required by Nov 30              │
│                                                                             │
│ Methods                                                                     │
│ ┌────────────────────────────────────────────────────────────────────────┐  │
│ │ TOTP Authenticator         [ Set up ]  [ Manage ]   Status: Not set    │  │
│ │ SMS Codes (+1 ***‑***‑1234) [ Manage ]               Status: Enabled   │  │
│ │ Passkeys (WebAuthn)        [ Set up ]               Status: Not set   │  │
│ └────────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│ Recovery Codes (10 available)     [ View / Generate ]                       │
│ Trusted Devices (5)               [ Manage ]                                │
│                                                                             │
│ Org Policy                                                               i  │
│ Your organization requires MFA. Complete setup to continue using all       │
│ features.  [ Go to MFA setup ]                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Mobile View (375–430px):**
```
┌──────────────────────────┐
│ ← Security & MFA         │
├──────────────────────────┤
│ MFA: Enabled  (Policy: Required) │
│ Methods                      │
│ • TOTP Authenticator  [Set up] │
│ • SMS (+1 ***‑***‑1234) [Manage] │
│ • Passkeys           [Set up]  │
│ Recovery Codes  (10) [View]    │
│ Trusted Devices  (5)  [Manage] │
│ [ Go to MFA setup ]            │
└──────────────────────────┘
```

---

#### Screen 2 — TOTP Setup (Step 1: QR + Secret)

```
┌──────────────────────────────────────────────┐
│ Set up TOTP Authenticator                    │
├──────────────────────────────────────────────┤
│ 1. Scan this QR with your authenticator app  │
│ [ █████████████ QR CODE █████████████ ]      │
│ Issuer: Skillsier   Account: user@example…   │
│ Can't scan? Use this key:  JBSWY3DPEHPK3PXP  │
│ [ Copy ]  [ Show as QR ]                     │
│ [ Continue ]                                 │
└──────────────────────────────────────────────┘
```

#### Screen 3 — TOTP Setup (Step 2: Verify Code)

```
┌──────────────────────────────────────────────┐
│ Enter the 6‑digit code from your app         │
│ Code: [ _ ][ _ ][ _ ][ _ ][ _ ][ _ ]         │
│ [ Verify ]     Didn’t work? [ Show tips ]    │
└──────────────────────────────────────────────┘
```

#### Screen 4 — SMS Setup (Phone → Code)

```
┌──────────────────────────────────────────────┐
│ Set up SMS Codes                             │
├──────────────────────────────────────────────┤
│ Phone number *  [+1] [ (___) ___‑____ ]      │
│ [ Send code ]                                 │
│                                                │
│ Code: [ _ ][ _ ][ _ ][ _ ][ _ ][ _ ]          │
│ [ Verify ]    [Resend 28s]                     │
└──────────────────────────────────────────────┘
```

#### Screen 5 — Passkey Enrollment

```
┌──────────────────────────────────────────────┐
│ Add a passkey                                │
├──────────────────────────────────────────────┤
│ This will use your device’s built‑in auth.   │
│ [ Continue ]                                 │
│                                              │
│ Nickname  [ My Laptop Passkey ]              │
└──────────────────────────────────────────────┘
```

#### Screen 6 — Recovery Codes

```
┌──────────────────────────────────────────────┐
│ Recovery Codes (Save these now)              │
├──────────────────────────────────────────────┤
│ • 82K4‑Z7MP • KDQ9‑L8RX • … (×10)            │
│ [ Download ]  [ Print ]  [ I saved these ]    │
└──────────────────────────────────────────────┘
```

#### Screen 7 — Disable Method (Danger Confirm)

```
┌──────────────────────────────────────────────┐
│ Disable TOTP?                                │
├──────────────────────────────────────────────┤
│ You must keep at least one MFA method.       │
│ Confirm with your password or a code.        │
│                                              │
│ Password [••••••]   or   Code [ _ _ _ _ _ _ ]│
│ [ Cancel ]                      [ Disable ]   │
└──────────────────────────────────────────────┘
```

#### Screen 8 — Trusted Devices (Revoke)

```
┌──────────────────────────────────────────────┐
│ Trusted Devices                              │
├──────────────────────────────────────────────┤
│ • Chrome on Mac — 197.22.4.18 — Last seen 2d │
│   [ Revoke ]                                 │
│ • iPhone 14 — LTE — Last seen 3h             │
│   [ Revoke ]                                 │
│                                              │
│ [ Revoke all ]                               │
└──────────────────────────────────────────────┘
```

#### Screen 9 — Policy Interstitial

```
┌──────────────────────────────────────────────┐
│ Your organization requires MFA               │
├──────────────────────────────────────────────┤
│ Complete setup to continue using all features│
│ [ Set up TOTP ] [ Set up SMS ] [ Add passkey ]│
└──────────────────────────────────────────────┘
```

---

## User Stories — Complete

> Abbrev: **W**=Web, **M**=Mobile.  
> Each story follows “As a … I want … so that …” and includes **components, routes, APIs**, plus **≥5 Happy** and **≥10 Bad** scenarios.

### Story 1 — View Security & MFA Overview

**Story**
```
As a Signed‑in User
I want to see my MFA status, available methods, recovery codes, and devices
So that I can understand and manage my account security
```

**Components & Routes**
- **W:** `(settings)/(security)/mfa/page.tsx` → `MfaOverview`, `Card`, `List`, `Badge`, `LinkButton`
- **M:** `…/security/mfa.tsx`
- **APIs:** `GET /v1/auth/mfa/methods`, `GET /v1/auth/trusted-devices`, `GET /v1/auth/mfa/policy`

**Acceptance Criteria — ✅ Happy**
1. **AC1.1** Overview loads within 1.0s (p95) and shows methods with status.
2. **AC1.2** Recovery codes status shows count or “Not generated”.
3. **AC1.3** Trusted devices shows count and quick link to manage page.
4. **AC1.4** Policy banner appears if `required=true` with due date.
5. **AC1.5** “Manage/Set up” buttons route to respective wizards.

**Acceptance Criteria — ❌ Bad**
1. **AC1.6** API 5xx → non‑blocking banner + Retry; partial data still displayed.
2. **AC1.7** Unauthorized (401/419) → redirect to re‑auth then back here.
3. **AC1.8** Empty methods array → shows clear empty state with guidance.
4. **AC1.9** Device list fetch fails → hides device card, shows link to retry.
5. **AC1.10** Policy fetch fails → assume not required; log warn event.
6. **AC1.11** Slow network (>2s) → skeleton remains, cancel lets user navigate.
7. **AC1.12** i18n key missing → English fallback; key logged.
8. **AC1.13** Screen reader cannot reach “Manage” → add aria‑labels.
9. **AC1.14** Bad cache shows stale status → refetch on focus re‑enters.
10. **AC1.15** Data mismatch across calls → reconcile by timestamp, show latest.

---

### Story 2 — Enroll TOTP Authenticator

**Story**
```
As a Security‑minded User
I want to pair a TOTP authenticator app and verify a code
So that I can use app‑based codes as a second factor
```

**Components & Routes**
- **W:** `(settings)/(security)/mfa/setup/totp/page.tsx` → `TotpSetupWizard`, `QRCode`, `CodeInput`
- **M:** `…/mfa-setup-totp.tsx` (mobile: show secret key prominently)
- **APIs:** `POST /v1/auth/mfa/totp/enroll/start`, `POST /v1/auth/mfa/totp/enroll/verify`

**Acceptance Criteria — ✅ Happy**
1. **AC2.1** Start call returns issuer, account, secret, otpauth URL; QR renders.
2. **AC2.2** “Can’t scan?” reveals secret key and copy button.
3. **AC2.3** Enter correct 6‑digit code → success; TOTP marked enabled.
4. **AC2.4** User prompted to generate recovery codes if none exist.
5. **AC2.5** Overview reflects TOTP “Enabled” immediately after success.

**Acceptance Criteria — ❌ Bad**
1. **AC2.6** Wrong code → inline error; after N attempts cool‑down starts.
2. **AC2.7** Time drift → guidance to sync device time; allow ±1 window.
3. **AC2.8** Start call fails → retry with backoff; wizard can restart safely.
4. **AC2.9** QR render blocked (image CSP) → fallback to text secret.
5. **AC2.10** Secret leaked into logs → test catches and blocks build.
6. **AC2.11** Double‑submit verify → idempotent; one success, one ignored.
7. **AC2.12** Mobile cannot scan itself → prominent “Copy secret” and tips.
8. **AC2.13** User navigates away mid‑flow → safe to resume; secret not persisted.
9. **AC2.14** Multiple TOTP enrollments exceed limit → clear error with limit.
10. **AC2.15** i18n numbers (Arabic‑Indic) → normalize to ASCII digits.

---

### Story 3 — Enroll SMS Codes

**Story**
```
As a User without an authenticator app
I want to receive one‑time codes by SMS
So that I can satisfy MFA with my phone number
```

**Components & Routes**
- **W:** `(settings)/(security)/mfa/setup/sms/page.tsx` → `SmsSetupWizard`, `PhoneInput`, `CodeInput`
- **M:** `…/mfa-setup-sms.tsx`
- **APIs:** `POST /v1/auth/mfa/sms/enroll/start`, `POST /v1/auth/mfa/sms/enroll/verify`, `POST /v1/auth/mfa/sms/enroll/resend`

**Acceptance Criteria — ✅ Happy**
1. **AC3.1** Phone input validates E.164; country picker sets format.
2. **AC3.2** Start returns `challengeId`; SMS arrives within typical SLA.
3. **AC3.3** Correct code → success; SMS listed as enabled.
4. **AC3.4** Resend disabled until cooldown; countdown shown & accessible.
5. **AC3.5** Masked phone displayed (+1 ***‑***‑1234) in confirmation.

**Acceptance Criteria — ❌ Bad**
1. **AC3.6** Invalid phone → inline error; prevents start.
2. **AC3.7** Wrong/expired code → error; after N tries restart challenge.
3. **AC3.8** Delivery delays → guidance & option to switch to TOTP.
4. **AC3.9** Carrier blocks short codes → show long code sender fallback.
5. **AC3.10** Roaming/number recycling risks → warning & confirmation.
6. **AC3.11** Resend spam → rate‑limited; explanatory tooltip.
7. **AC3.12** ChallengeId missing → restart flow; no PII leaked.
8. **AC3.13** Dual‑SIM misroute → allow resend with changed SIM prompt.
9. **AC3.14** Screen reader skips resend → aria‑live polite message added.
10. **AC3.15** Local numerals → normalize; still accept ASCII only for code.

---

### Story 4 — Enroll Passkey (WebAuthn/FIDO2)

**Story**
```
As a Security‑savvy User
I want to enroll a passkey (security key or platform)
So that I can use phishing‑resistant MFA
```

**Components & Routes**
- **W:** `(settings)/(security)/mfa/setup/passkey/page.tsx` → `PasskeyEnrollPanel`, `Button`, `Input (nickname)`
- **M:** `…/mfa-setup-passkey.tsx` (platform support dependent)
- **APIs:** `POST /v1/auth/webauthn/register/start`, `POST /v1/auth/webauthn/register/finish`

**Acceptance Criteria — ✅ Happy**
1. **AC4.1** Start returns creation options; browser prompt appears.
2. **AC4.2** Finish stores credential; user can set nickname.
3. **AC4.3** Credential visible in overview; test sign‑in passes.
4. **AC4.4** Multiple credentials supported up to limit; list shows all.
5. **AC4.5** Analytics captures enroll success (no PII).

**Acceptance Criteria — ❌ Bad**
1. **AC4.6** Unsupported browser → button disabled + tooltip guidance.
2. **AC4.7** User cancels prompt → neutral cancel state; no error.
3. **AC4.8** Origin/RP mismatch → hard error; guidance to use official domain.
4. **AC4.9** Platform authenticator locked → prompt to unlock/try again.
5. **AC4.10** Finish fails (attestation invalid) → explain and recommend retry.
6. **AC4.11** Duplicate credentialId → backend idempotency prevents dupes.
7. **AC4.12** Nickname validation fails → inline error, retains value.
8. **AC4.13** Security policy forbids (org) → show policy reason and alternatives.
9. **AC4.14** Credential created but list fetch fails → soft error; still enabled.
10. **AC4.15** Device time skew breaks challenge → re‑start flow gracefully.

---

### Story 5 — Generate & Manage Recovery Codes

**Story**
```
As a Safety‑conscious User
I want to generate and securely store recovery codes
So that I can sign in when I lose my device
```

**Components & Routes**
- **W:** `(settings)/(security)/mfa/recovery-codes/page.tsx` → `RecoveryCodesPanel`
- **M:** `…/mfa-recovery-codes.tsx`
- **APIs:** `POST /v1/auth/mfa/recovery/generate`

**Acceptance Criteria — ✅ Happy**
1. **AC5.1** Generating shows 10 unique codes; previous set invalidated.
2. **AC5.2** Download (CSV) and Print options available; “I saved these” confirm.
3. **AC5.3** Overview shows “10 available” after generate.
4. **AC5.4** Analytics logs generate event without codes content.
5. **AC5.5** Accessibility: codes selectable and readable by screen readers.

**Acceptance Criteria — ❌ Bad**
1. **AC5.6** Regenerate too frequently → rate limit + cooldown.
2. **AC5.7** Print blocked by browser → message with copy alternative.
3. **AC5.8** Clipboard blocked → fallback instructions.
4. **AC5.9** Codes leak in logs → blocked by test; scrub middleware.
5. **AC5.10** Concurrent generate from two tabs → last set wins; warning.
6. **AC5.11** Lost focus causes accidental dismissal → require explicit confirm.
7. **AC5.12** Download fails → retry; offer copy‑to‑clipboard chunked.
8. **AC5.13** Non‑monospace wraps badly → fixed‑width style applied.
9. **AC5.14** Localization breaks CSV delimiter → locale‑safe exporter.
10. **AC5.15** Low‑vision users struggle → high‑contrast theme supported.

---

### Story 6 — Disable/Remove an MFA Method (with Safeguards)

**Story**
```
As a Careful User
I want to disable an MFA method safely
So that I never lock myself out
```

**Components & Routes**
- **W:** `(settings)/(security)/mfa/disable/[method]/page.tsx` → `DisableMethodDialog`
- **M:** Confirm dialog in respective screen
- **APIs:** `POST /v1/auth/mfa/disable`

**Acceptance Criteria — ✅ Happy**
1. **AC6.1** If more than one method enabled → disable allowed after re‑auth.
2. **AC6.2** If only one method exists → UI blocks with guidance to add another.
3. **AC6.3** Post‑disable, overview updates instantly.
4. **AC6.4** Audit log entry created; email notification sent.
5. **AC6.5** Recovery codes remain valid; user advised to rotate if leaked.

**Acceptance Criteria — ❌ Bad**
1. **AC6.6** Re‑auth fails → disable blocked with retry option.
2. **AC6.7** Backend rejects due to policy → error shows policy citation.
3. **AC6.8** Race condition (another tab disables same) → idempotent success.
4. **AC6.9** User has no second factor and insists → requires admin flow (blocked).
5. **AC6.10** Network/5xx → safe retry; no partial state.
6. **AC6.11** Email notification fails → still disabled; event queued.
7. **AC6.12** Accessibility: confirm modal traps focus; ensure proper traps.
8. **AC6.13** Mobile back gesture dismisses without confirm → ask again.
9. **AC6.14** CSRF invalid → hard refresh and retry.
10. **AC6.15** i18n missing strings → English fallback + log.

---

### Story 7 — Manage Trusted Devices

**Story**
```
As a User
I want to review and revoke trusted devices
So that I can keep my account secure
```

**Components & Routes**
- **W:** `(settings)/(security)/mfa/devices/page.tsx` → `TrustedDevicesTable`
- **M:** `…/mfa-devices.tsx`
- **APIs:** `GET /v1/auth/trusted-devices`, `DELETE /v1/auth/trusted-devices/{id}`

**Acceptance Criteria — ✅ Happy**
1. **AC7.1** List shows UA, IP, created, last seen, approximate location.
2. **AC7.2** Revoke one → requires confirm → device cookie invalidated.
3. **AC7.3** Revoke all → logs out all remembered devices.
4. **AC7.4** Overview count updates after revoke.
5. **AC7.5** Email notification summarizes changes.

**Acceptance Criteria — ❌ Bad**
1. **AC7.6** Revoke call fails → error; state remains unchanged.
2. **AC7.7** Device already revoked → idempotent success.
3. **AC7.8** GeoIP unavailable → label as “Unknown location”.
4. **AC7.9** Large list pagination → infinite scroll with skeleton.
5. **AC7.10** Offline → queue action and retry when online.
6. **AC7.11** Session expired → re‑auth then resume action.
7. **AC7.12** Conflicting actions (revoke + logout) → serialize safely.
8. **AC7.13** Accessibility: table labels and buttons announced.
9. **AC7.14** Rate limit on revoke all → throttle & feedback.
10. **AC7.15** Browser blocks cookies → note that trust device may not persist.

---

### Story 8 — MFA Policy Enforcement Interstitial

**Story**
```
As an Organization Admin
I want MFA to be required for members
So that our workspace remains secure
```

**Components & Routes**
- **W:** `(settings)/(security)/mfa/policy/page.tsx` → `PolicyInterstitial`
- **M:** `…/mfa-policy.tsx`
- **APIs:** `GET /v1/auth/mfa/policy`

**Acceptance Criteria — ✅ Happy**
1. **AC8.1** If `required=true`, interstitial blocks access to sensitive routes.
2. **AC8.2** Progress shows completed vs remaining methods.
3. **AC8.3** On completion, user redirected to original destination.
4. **AC8.4** Admin policy changes reflected within 60s (refetch/onFocus).
5. **AC8.5** Analytics records compliance completion.

**Acceptance Criteria — ❌ Bad**
1. **AC8.6** Policy fetch fails → soft warning; allow limited access.
2. **AC8.7** User tries to bypass via URL → redirect back to interstitial.
3. **AC8.8** Timezone confusion on due date → display absolute UTC + local.
4. **AC8.9** Disabled JS → server enforces via middleware (SSR redirect).
5. **AC8.10** Lost network mid‑setup → preserve step and resume.
6. **AC8.11** Admin removes requirement mid‑flow → remove banner and continue.
7. **AC8.12** Multiple tabs compete → first completion wins; others sync.
8. **AC8.13** A11y: interstitial focus trapped; escape and navigation keys work.
9. **AC8.14** SSO bypass suspected → still enforce MFA post‑SSO hook.
10. **AC8.15** Privacy: banner reveals too much about policy → minimal copy.

---

### Story 9 — Lost Device & Recovery Path

**Story**
```
As a Locked‑out User
I want clear instructions to recover access using recovery codes
So that I can sign in and re‑enroll MFA securely
```

**Components & Routes**
- **W/M:** Links from overview to AUTH‑2 recovery‑code login
- **APIs:** N/A in this screen; relies on AUTH‑2 endpoints

**Acceptance Criteria — ✅ Happy**
1. **AC9.1** Overview explains recovery code usage and where to enter it.
2. **AC9.2** After using a recovery code, prompt to regenerate a new set.
3. **AC9.3** Suggest re‑enrolling primary method (TOTP/Passkey).
4. **AC9.4** Security email alerts sent after recovery usage.
5. **AC9.5** Analytics records recovery usage event (no code value).

**Acceptance Criteria — ❌ Bad**
1. **AC9.6** No recovery codes available → link to generate (requires any factor).
2. **AC9.7** Confusing wording → doc copy lint fails; requires rewrite.
3. **AC9.8** Link loops back incorrectly → fix redirect to AUTH‑2 exact route.
4. **AC9.9** Mobile deep link missing → fallback to web URL with QR.
5. **AC9.10** Support escalation unclear → clear “Contact support” path.
6. **AC9.11** A11y: long paragraphs → readable line length, headings.
7. **AC9.12** Localization issues in guidance → keys reviewed.
8. **AC9.13** Telemetry includes PII → scrub context.
9. **AC9.14** External doc link broken → show inline guidance instead.
10. **AC9.15** User confuses recovery vs backup codes → glossary tooltip added.

---

### Story 10 — Analytics & Telemetry for MFA Management

**Story**
```
As a Product Team
I want complete instrumentation of MFA management
So that I can monitor adoption and reliability
```

**Components & Routes**
- Emitted across pages; no dedicated route

**Events (examples)**
- `security.mfa.view_overview`
- `security.mfa.totp.start` / `.verify.success` / `.verify.error`
- `security.mfa.sms.start` / `.verify.success` / `.verify.error` / `.resend`
- `security.mfa.passkey.start` / `.finish.success` / `.finish.error`
- `security.mfa.recovery.generate`
- `security.mfa.device.revoke_one` / `.revoke_all`
- `security.mfa.policy.view` / `.complete`

**Acceptance Criteria — ✅ Happy**
1. **AC10.1** Events are idempotent and debounced; no duplicates on rerenders.
2. **AC10.2** Payloads contain method identifiers but never secrets/codes.
3. **AC10.3** SDK failures are non‑blocking and queued for retry.
4. **AC10.4** Versioned event names; schema validation in CI.
5. **AC10.5** Dashboard shows conversion from “start” → “success” per method.

**Acceptance Criteria — ❌ Bad**
1. **AC10.6** PII present in payload → build fails; linter catches.
2. **AC10.7** Missed events due to route transitions → use visibility/focus hooks.
3. **AC10.8** Time skew across clients → server timestamp added.
4. **AC10.9** Over‑instrumentation hurts performance → budget enforced.
5. **AC10.10** Ad‑blockers drop beacons → offline queue with backoff.

---

## Non‑Functional Requirements (AUTH‑4 scope)

- **Performance:** Overview fetch < 1000ms (p95); setup verify < 600ms (p95)
- **Security:** CSRF (web), rate limits, WebAuthn attestation validation, TOTP drift window, SMS challenge TTL, recovery codes hashing; Sentry
- **Privacy:** No secrets in logs/analytics; masked phone; minimal policy exposure
- **Compatibility:** Web (Chrome/Firefox/Safari/Edge), iOS 14+, Android 10+; passkeys when supported
- **Testing:** Unit tests (schemas/hooks/utils), e2e (TOTP/SMS/Passkey/Recovery), a11y checks (axe), policy enforcement tests

---

**End of AUTH‑4 Enhancements (Complete).**
