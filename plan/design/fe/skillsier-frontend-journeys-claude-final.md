# Section A: AUTHENTICATION & ONBOARDING

## Section A.1 Authentication (AUTH)

### Journey: AUTH-1 (User Registration)

**Goal:** Allow new users to create an account on Skillsier platform

**Actor:** Anonymous User

**Trigger:** User clicks "Sign Up" from landing page or login page

**Frontend Routes:**
- Web: `/auth/register`
- Mobile: `app/(auth)/register.tsx`

**Backend Services:**
- **Primary:** users-be/user (`POST /v1/users/register`)
- **Secondary:** Keycloak OAuth2 (authentication provider)
- **Secondary:** users-be/verification (`POST /v1/users/{id}/verify-email/send`)

**Key Components:**
- `apps/web/app/(auth)/register/page.tsx` - Registration form
- `apps/mobile/app/(auth)/register.tsx` - Mobile registration
- `packages/lib/src/features/auth/hooks/use-register.ts` - Registration hook
- `packages/lib/src/features/auth/components/RegisterForm.tsx` - Form component
- `packages/ui/src/components/forms/PasswordStrengthMeter.tsx` - Password validation UI

---

#### Flow Steps

**Step 1: Navigate to Registration Page**
- User clicks "Sign Up" button from:
  - Landing page (`/(public)/page.tsx`)
  - Login page (`/auth/login`)
  - Job detail page (when trying to apply)
  - Freelancer profile page (when trying to contact)
- Redirect to `/auth/register`

**Step 2: Choose Account Type**
- **Display:**
  - Title: "Join Skillsier"
  - Subtitle: "I want to..."
  - Two large cards:
    - **Freelancer Card:**
      - Icon: Briefcase
      - Title: "Work as a Freelancer"
      - Description: "Find projects and build your business"
      - Button: "Join as Freelancer"
    - **Client Card:**
      - Icon: Search
      - Title: "Hire Talent"
      - Description: "Find the perfect freelancer for your project"
      - Button: "Join as Client"
  - Link: "Already have an account? Log in"
- User selects account type (sets `userType` state)

**Step 3: Registration Form Display**
- Form fields:
  - **First Name:**
    - Input type: text
    - Required: Yes
    - Validation: 2-50 characters, letters only
    - Error: "First name is required"
  - **Last Name:**
    - Input type: text
    - Required: Yes
    - Validation: 2-50 characters, letters only
    - Error: "Last name is required"
  - **Email:**
    - Input type: email
    - Required: Yes
    - Validation: Valid email format
    - Async validation: Check if email exists
    - Error: "Email is already registered"
  - **Password:**
    - Input type: password
    - Required: Yes
    - Validation:
      - Minimum 8 characters
      - At least 1 uppercase letter
      - At least 1 lowercase letter
      - At least 1 number
      - At least 1 special character (!@#$%^&*)
    - Show password toggle (eye icon)
    - Password strength meter:
      - Weak (red) - Less than 50% criteria met
      - Medium (yellow) - 50-75% criteria met
      - Strong (green) - 100% criteria met
    - Real-time feedback as user types
  - **Confirm Password:**
    - Input type: password
    - Required: Yes
    - Validation: Must match password
    - Error: "Passwords do not match"
  - **Terms Checkbox:**
    - Checkbox: "I agree to the Terms of Service and Privacy Policy"
    - Links open in new tab:
      - Terms of Service → `/legal/terms`
      - Privacy Policy → `/legal/privacy`
    - Required: Yes
    - Error: "You must accept the terms to continue"
  - **Submit Button:**
    - Text: "Create Account"
    - Disabled if: Form invalid or submitting
    - Loading state: Shows spinner when submitting

**Step 4: Social Registration Options (Alternative via Keycloak)**
- Separator: "or sign up with"
- Social button:
  - **Google (Gmail):**
    - Icon: Google logo
    - Text: "Continue with Google"
    - Action: Initiate OAuth2 flow with Keycloak
    - Flow:
      1. Frontend redirects to Keycloak authorization endpoint:
         ```
         GET https://keycloak.skillsier.com/realms/skillsier/protocol/openid-connect/auth
         ?client_id=skillsier-web
         &redirect_uri=https://skillsier.com/auth/callback
         &response_type=code
         &scope=openid email profile
         &state={random_state}
         &kc_idp_hint=google
         ```
      2. User redirected to Google OAuth consent screen
      3. User grants permissions to Skillsier
      4. Google redirects back to Keycloak with authorization code
      5. Keycloak exchanges code for Google access token
      6. Keycloak retrieves user info from Google
      7. Keycloak redirects to skillsier.com/auth/callback with authorization code
      8. Frontend exchanges code for Keycloak JWT tokens
      9. If first-time Google user:
         - Keycloak creates new user automatically
         - users-be receives Keycloak event webhook
         - users-be creates user record with `keycloakId`
         - Email auto-verified (trusted from Google)
         - User redirected to account type selection: `/onboarding/select-type`
      10. If existing Google user:
          - Keycloak returns existing user tokens
          - User logged in directly to dashboard

**Note:** Gmail OAuth is the only social login option currently configured in Keycloak

**Step 5: Form Submission (Email/Password via Keycloak)**
- User clicks "Create Account"
- Frontend validation:
  - Check all required fields filled
  - Validate email format
  - Validate password strength
  - Confirm passwords match
  - Check terms accepted
- If validation fails:
  - Show inline error messages
  - Focus first error field
  - Disable submit button
- If validation passes:
  - Show loading spinner on button
  - Disable all form inputs
- **Keycloak Registration Flow:**
  1. **Create Keycloak User:**
     - API Call to Keycloak Admin API (via users-be proxy):
     ```typescript
     POST /v1/auth/register
     Headers: {
       Content-Type: application/json
     }
     Body: {
       firstName: string,
       lastName: string,
       email: string,
       password: string,
       userType: 'CLIENT' | 'FREELANCER',
       agreedToTerms: boolean,
       agreedToTermsAt: ISO8601 timestamp
     }
     ```
  2. **users-be handles:**
     - Creates user in Keycloak
     - Receives `keycloakId` from Keycloak
     - Creates user record in users-be database with:
       - `keycloakId`: Reference to Keycloak user
       - `email`: User email
       - `emailVerified`: false (initially)
       - `userType`: CLIENT or FREELANCER
     - Returns JWT tokens from Keycloak OAuth2 flow
- Loading state: "Creating your account..."

**Step 6: Handle Registration Response**

**Success Response (201 Created):**
```json
{
  "userId": "uuid",
  "keycloakId": "uuid",
  "email": "user@example.com",
  "firstName": "John",
  "lastName": "Doe",
  "userType": "FREELANCER",
  "emailVerified": false,
  "emailVerificationRequired": true,
  "emailVerificationSentAt": "2025-11-07T10:00:00Z",
  "tokens": {
    "accessToken": "keycloak_jwt_access_token",
    "refreshToken": "keycloak_jwt_refresh_token",
    "idToken": "keycloak_id_token",
    "tokenType": "Bearer",
    "expiresIn": 3600,
    "refreshExpiresIn": 86400
  }
}
```
- **Token Storage:**
  - `accessToken`: Stored in memory (AuthContext state)
  - `refreshToken`: Stored in HttpOnly secure cookie (handled by backend)
  - `idToken`: Stored in memory (contains user claims)
  - All tokens are Keycloak-issued JWTs
- **Token Usage:**
  - Access token sent in Authorization header: `Bearer {accessToken}`
  - Access token contains claims:
    ```json
    {
      "sub": "keycloak_user_id",
      "email": "user@example.com",
      "email_verified": false,
      "name": "John Doe",
      "preferred_username": "user@example.com",
      "given_name": "John",
      "family_name": "Doe",
      "realm_access": {
        "roles": ["user", "freelancer"]
      }
    }
    ```
- Update auth context with user data
- Trigger background events:
  - Keycloak publishes user creation event
  - users-be listens to Keycloak webhook
  - users-be creates user record with `keycloakId`
  - users-be publishes `user.created.v1` event
  - Verification email queued (if email not verified)

**Error Responses:**

**400 Bad Request - Validation Error:**
```json
{
  "error": "VALIDATION_ERROR",
  "message": "Invalid input",
  "details": [
    {
      "field": "email",
      "message": "Email is already registered"
    }
  ]
}
```
- Display inline error under relevant field
- Re-enable form for editing

**409 Conflict - Email Exists:**
```json
{
  "error": "EMAIL_EXISTS",
  "message": "An account with this email already exists"
}
```
- Show error toast: "This email is already registered. Try logging in?"
- Offer "Login instead" button
- Clear password field for security

**429 Too Many Requests:**
```json
{
  "error": "RATE_LIMIT_EXCEEDED",
  "message": "Too many registration attempts. Please try again in 15 minutes."
}
```
- Show error toast with retry time
- Disable form for 15 minutes
- Show countdown timer

**500 Internal Server Error:**
```json
{
  "error": "INTERNAL_ERROR",
  "message": "Something went wrong. Please try again."
}
```
- Show error toast: "Unable to create account. Please try again."
- Log error to monitoring service
- Re-enable form

**Step 7: Email Verification Prompt**
- On successful registration:
  - Navigate to `/auth/verify-email`
  - Display:
    - Icon: Email envelope
    - Title: "Verify your email"
    - Message: "We sent a verification link to **{email}**"
    - Instructions: "Click the link in the email to verify your account"
    - Timer: "Email expires in 15 minutes"
    - Actions:
      - "Resend Email" button (disabled for 60 seconds after send)
      - "Change Email" link (returns to registration with email pre-filled)
      - "Open Email App" button (deep link to email app on mobile)
- Background:
  - Email sent via users-be to notification service
  - Email contains verification link: `https://skillsier.com/auth/verify-email?token={verification_token}`
  - Token valid for 15 minutes
  - User record marked as `emailVerified: false`

**Step 8: Email Verification (User Clicks Link)**
- User clicks link in email
- Navigate to `/auth/verify-email?token={token}`
- API Call:
  ```typescript
  POST /v1/users/verify-email
  Body: {
    token: string
  }
  ```
- Backend validates token:
  - Check token not expired
  - Check token not already used
  - Mark email as verified
  - Publish `user.verified.v1` event

**Success Response:**
```json
{
  "success": true,
  "message": "Email verified successfully"
}
```
- Display success screen:
  - Icon: Green checkmark
  - Title: "Email Verified!"
  - Message: "Your account is ready to use"
  - Button: "Continue to Dashboard"
- Auto-redirect after 3 seconds to:
  - Freelancer: `/onboarding/profile-setup` (onboarding flow)
  - Client: `/dashboard` (main dashboard)

**Error Responses:**

**400 - Invalid Token:**
- Display error:
  - Icon: Red X
  - Title: "Invalid Verification Link"
  - Message: "This link is invalid or has expired"
  - Button: "Resend Verification Email"

**410 - Token Expired:**
- Display error:
  - Icon: Clock
  - Title: "Link Expired"
  - Message: "Verification links expire after 15 minutes"
  - Button: "Send New Link"

**409 - Already Verified:**
- Display message:
  - Icon: Green checkmark
  - Title: "Already Verified"
  - Message: "Your email is already verified"
  - Button: "Go to Dashboard"

**Step 9: Resend Verification Email**
- User clicks "Resend Email" button
- Button disabled for 60 seconds after send (rate limiting)
- Show countdown: "Resend available in 59s..."
- API Call:
  ```typescript
  POST /v1/users/verify-email/resend
  Headers: {
    Authorization: Bearer {accessToken}
  }
  ```
- Backend:
  - Validates user is authenticated
  - Checks email not already verified
  - Generates new verification token
  - Invalidates old token
  - Sends new email via notification service
- Success:
  - Toast: "Verification email sent!"
  - Reset 60-second timer
- Error (429 Too Many Requests):
  - Toast: "Too many requests. Try again in 5 minutes."

**Step 10: OAuth Callback Handler**
- Route: `/auth/callback`
- Query params: `?code={auth_code}&state={state}`
- Flow:
  1. Validate `state` parameter matches stored state (CSRF protection)
  2. Exchange authorization code for tokens:
     ```typescript
     POST /v1/auth/callback
     Body: {
       code: string,
       state: string
     }
     ```
  3. Backend:
     - Exchanges code with Keycloak token endpoint
     - Receives access_token, refresh_token, id_token
     - Validates tokens
     - Checks if user exists in users-be database
     - If new OAuth user:
       - Create user record with `keycloakId`
       - Set `emailVerified: true` (trusted OAuth provider)
       - Redirect to `/onboarding/select-type`
     - If existing user:
       - Update last login timestamp
       - Redirect to `/dashboard`
  4. Store tokens in frontend
  5. Navigate based on account state

---

#### Edge Cases & Alternate Flows

**C1: Network Error During Registration**
- Registration API call fails due to network issue
- Display error toast: "Network error. Please check your connection."
- Re-enable form
- Preserve form data (don't clear fields)
- User can retry submission

**C2: Duplicate Registration Attempt**
- User already registered but forgot
- Clicks "Sign Up" again with same email
- Backend returns 409 Conflict
- Display: "This email is already registered"
- Show dialog:
  - Message: "An account with this email already exists"
  - Options:
    - "Login" button → Navigate to `/auth/login`
    - "Forgot Password?" link → Navigate to `/auth/forgot-password`
    - "Cancel" button → Close dialog

**C3: Email Verification Token Expired**
- User clicks verification link after 15+ minutes
- Backend returns 410 Gone
- Display:
  - Title: "Link Expired"
  - Message: "Verification links expire after 15 minutes for security"
  - Button: "Send New Link"
- User clicks button
- New verification email sent
- Toast: "New verification email sent!"

**C4: Weak Password Attempt**
- User enters password that doesn't meet criteria
- Real-time feedback shows:
  - Password strength: "Weak"
  - Missing criteria highlighted in red:
    - ❌ At least 8 characters
    - ❌ 1 uppercase letter
    - ✅ 1 lowercase letter
    - ✅ 1 number
    - ❌ 1 special character
- Submit button disabled
- Tooltip: "Password must meet all requirements"

**C5: Terms Not Accepted**
- User tries to submit without checking terms checkbox
- Form validation fails
- Error message under checkbox: "You must accept the terms to continue"
- Checkbox highlighted in red
- Focus moved to checkbox

**C6: OAuth Account Linking Conflict**
- User registers with email/password: "john@example.com"
- Later tries OAuth login with same email
- Keycloak detects email conflict
- Display dialog:
  - Title: "Account Linking Required"
  - Message: "An account with this email already exists. Link your Google account?"
  - Options:
    - "Link Accounts" → Prompts for password to verify
    - "Use Different Email" → Return to registration
    - "Cancel" → Return to login page
- If user links:
  - Prompt for current password
  - Verify password with Keycloak
  - Link OAuth identity to existing account
  - Success: "Accounts linked successfully!"

**C7: Invalid Email Format**
- User enters malformed email: "userexample.com"
- Real-time validation shows error: "Please enter a valid email"
- Submit button disabled
- Email field highlighted in red

**C8: Special Characters in Name**
- User enters name with numbers: "John123"
- Validation error: "Name can only contain letters"
- Field highlighted in red
- Focus remains on field

**C9: Keycloak Service Down**
- Registration request fails due to Keycloak unavailability
- Backend returns 503 Service Unavailable
- Display error:
  - Toast: "Authentication service temporarily unavailable"
  - Message: "We're experiencing technical difficulties. Please try again in a few minutes."
  - Hide form inputs
  - Show retry button
- Log incident to monitoring
- Show status page link if outage persists

**C10: Browser Back Button During Registration**
- User clicks back button during registration
- If form has unsaved data:
  - Browser prompt: "Leave site? Changes you made may not be saved"
  - User can choose: "Leave" or "Stay"
- If user leaves: Data lost (expected behavior)
- If user stays: Form data preserved

**C11: Password Paste Disabled in Confirm Field**
- User tries to paste password in "Confirm Password" field
- Paste disabled (security best practice)
- Tooltip on hover: "Please type your password to confirm"
- User must manually type password

**C12: Email Verification Spam Filter**
- Verification email caught by spam filter
- User doesn't receive email
- After 5 minutes, user clicks "Resend Email"
- Display tip: "Can't find the email? Check your spam folder"
- Provide support email: support@skillsier.com
- "Still having issues?" link to help center

**C13: Multi-Device Registration Flow**
- User starts registration on mobile
- Verification email received on desktop
- User clicks link on desktop browser
- Keycloak verifies token (device-agnostic)
- Desktop browser now logged in
- Mobile app still shows "Verify email" screen
- Mobile app polls verification status every 10 seconds
- On verification detected:
  - Mobile app updates UI to "Email Verified!"
  - Auto-navigate to onboarding

**C14: Rate Limiting on Registration**
- Too many registration attempts from same IP
- Backend returns 429 Too Many Requests
- Display:
  - Error toast: "Too many registration attempts"
  - Message: "Please wait 15 minutes before trying again"
  - Countdown timer: "Retry available in 14:59"
  - Disable all form inputs
- Log suspicious activity for fraud detection

**C15: Incomplete OAuth Flow**
- User clicks "Continue with Google"
- OAuth consent screen opens
- User closes window without consenting
- Frontend detects window close
- Show message: "Login canceled"
- Return to registration page
- Form data preserved

**C16: OAuth Email Mismatch**
- User starts registration with: "john@example.com"
- Clicks "Continue with Google"
- Logs in with different Google account: "john.work@company.com"
- Keycloak creates account with: "john.work@company.com"
- users-be creates record with OAuth email
- Success flow proceeds normally
- Note: Email from OAuth provider is source of truth

**C17: Accessibility - Screen Reader**
- Screen reader announces form labels clearly
- Validation errors read aloud immediately
- Password strength meter has aria-live region
- Submit button state announced: "disabled" or "enabled"
- Focus management: Error focuses first invalid field
- All interactive elements keyboard accessible

**C18: Mobile Native Biometric After Registration**
- User completes registration on mobile app
- After email verification
- Prompt: "Enable Face ID / Touch ID for quick login?"
- Options: "Enable" or "Skip"
- If enabled:
  - Store refresh token in secure keychain
  - Associate with biometric
  - Future logins: Biometric → Auto-login
- If skipped: Standard login flow required

---

#### Notifications

**Registration Success:**
- **Email:** "Welcome to Skillsier!" (HTML template)
  - Subject: "Welcome to Skillsier - Verify Your Email"
  - Body:
    - Greeting: "Hi {firstName},"
    - Message: "Welcome to Skillsier! Click the button below to verify your email"
    - CTA Button: "Verify Email" (links to verification URL)
    - Alternative: Plain text link below button
    - Expiry notice: "This link expires in 15 minutes"
    - Footer: Support email, social links

**Email Verified:**
- **In-app notification:** "Email verified! Your account is ready."
- **Email:** "Email Verified - Get Started" (optional)
  - Next steps based on user type:
    - Freelancer: "Complete your profile", "Add skills", "Browse jobs"
    - Client: "Post a job", "Browse talent", "Set up payment method"

**Failed Login Attempt After Registration:**
- **Email:** "Unusual Login Activity Detected"
  - Triggered if: Multiple failed login attempts after registration
  - Message: "We detected unusual login activity. If this wasn't you, secure your account."
  - CTA: "Secure My Account" → Password reset flow

---

#### Analytics

- `auth.registration.started` (user_type, registration_method)
- `auth.registration.completed` (user_type, registration_method, duration_seconds)
- `auth.registration.failed` (error_code, error_message, user_type)
- `auth.email_verification.sent` (user_id)
- `auth.email_verification.completed` (user_id, time_to_verify_minutes)
- `auth.email_verification.failed` (user_id, error_reason)
- `auth.oauth.started` (provider)
- `auth.oauth.completed` (provider, new_user)
- `auth.oauth.failed` (provider, error_code)
- `auth.password_strength.evaluated` (strength_score, user_type)
- `auth.terms.accepted` (user_id, terms_version)

---

#### Sources

**Frontend:**
- `apps/web/app/(auth)/register/page.tsx`
- `apps/mobile/app/(auth)/register.tsx`
- `packages/lib/src/features/auth/hooks/use-register.ts`
- `packages/lib/src/features/auth/components/RegisterForm.tsx`
- `packages/ui/src/components/forms/PasswordStrengthMeter.tsx`

**Backend:**
- `users-be/internal/domain/user/` - User entity and repository
- `users-be/internal/interfaces/http/v1/handlers/auth_handler.go` - Registration endpoints
- `users-be/internal/application/auth/service.go` - Registration business logic
- Keycloak Admin API - User creation
- Keycloak OAuth2 endpoints - Social login flows

**Database:**
- `users-be`: `users` table
  - Fields: id, keycloak_id, email, email_verified, first_name, last_name, user_type, created_at

**Events:**
- Published: `user.created.v1`, `user.verified.v1`
- Consumed: Keycloak user creation events (webhook)

---

### Journey: AUTH-2 (User Login)

**Goal:** Allow existing users to authenticate and access their account

**Actor:** Registered User (email not verified or verified)

**Trigger:** User clicks "Login" from any page

**Frontend Routes:**
- Web: `/auth/login`
- Mobile: `app/(auth)/login.tsx`

**Backend Services:**
- **Primary:** Keycloak OAuth2 (`POST /realms/skillsier/protocol/openid-connect/token`)
- **Secondary:** users-be/auth (`POST /v1/auth/login` - proxy to Keycloak)
- **Secondary:** users-be/security/sessions (`POST /v1/auth/sessions`)

---

#### Flow Steps

**Step 1: Navigate to Login Page**
- User clicks "Login" button from:
  - Landing page header
  - Registration page ("Already have an account?")
  - Protected route redirect (user tries to access without auth)
  - Job application prompt
- Redirect to `/auth/login`
- If accessed via redirect, store `returnUrl` in session storage

**Step 2: Login Form Display**
- **Header:**
  - Title: "Welcome Back"
  - Subtitle: "Login to your Skillsier account"
- **Form Fields:**
  - **Email / Username:**
    - Input type: email
    - Placeholder: "Email address"
    - Required: Yes
    - Autofocus: Yes
    - Autocomplete: "email"
  - **Password:**
    - Input type: password
    - Placeholder: "Password"
    - Required: Yes
    - Show/hide toggle (eye icon)
    - Autocomplete: "current-password"
  - **Remember Me:**
    - Checkbox: "Keep me logged in"
    - Default: Unchecked
    - Note: Extends refresh token expiry to 30 days
  - **Forgot Password Link:**
    - Text: "Forgot password?"
    - Navigate to: `/auth/forgot-password`
  - **Submit Button:**
    - Text: "Login"
    - Disabled if form invalid
    - Loading state: Shows spinner
- **Social Login:**
  - Separator: "or continue with"
  - **Google Button:**
    - Icon: Google logo
    - Text: "Continue with Google"
    - Action: Initiate Keycloak OAuth flow
- **Registration Link:**
  - Text: "Don't have an account? Sign up"
  - Navigate to: `/auth/register`

**Step 3: Form Submission (Email/Password via Keycloak)**
- User clicks "Login"
- Frontend validation:
  - Email/username not empty
  - Password not empty
- If validation passes:
  - Show loading spinner
  - Disable form inputs
- **Keycloak Login Flow:**
  - API Call (via users-be proxy):
    ```typescript
    POST /v1/auth/login
    Headers: {
      Content-Type: application/json
    }
    Body: {
      username: string,  // email or username
      password: string,
      rememberMe: boolean
    }
    ```
  - Backend forwards to Keycloak:
    ```typescript
    POST /realms/skillsier/protocol/openid-connect/token
    Body: {
      grant_type: "password",
      client_id: "skillsier-web",
      client_secret: "{client_secret}",
      username: "{email}",
      password: "{password}",
      scope: "openid email profile"
    }
    ```
  - Keycloak validates credentials against its user database
  - On success: Returns JWT tokens
  - On failure: Returns error response

**Step 4: Handle Login Response**

**Success Response (200 OK):**
```json
{
  "accessToken": "keycloak_jwt_access_token",
  "refreshToken": "keycloak_jwt_refresh_token",
  "idToken": "keycloak_id_token",
  "tokenType": "Bearer",
  "expiresIn": 3600,
  "refreshExpiresIn": 86400,  // 30 days if rememberMe: true
  "user": {
    "userId": "uuid",
    "keycloakId": "uuid",
    "email": "user@example.com",
    "firstName": "John",
    "lastName": "Doe",
    "userType": "FREELANCER",
    "emailVerified": true,
    "profileCompletionScore": 75
  }
}
```
- **Token Storage:**
  - Access token: Memory (AuthContext)
  - Refresh token: HttpOnly secure cookie
  - ID token: Memory
- **Update Auth State:**
  - Set user in AuthContext
  - Set authenticated: true
  - Store user preferences in localStorage (non-sensitive)
- **Session Tracking:**
  - Backend creates session record:
    ```typescript
    POST /v1/auth/sessions
    Body: {
      userId: uuid,
      sessionToken: hashed,
      ipAddress: string,
      userAgent: string,
      deviceId: string
    }
    ```
  - Session stored in `sessions` table
  - Session expiry tracked
- **Navigation:**
  - If `returnUrl` exists: Navigate to returnUrl
  - Else if email not verified: Navigate to `/auth/verify-email`
  - Else: Navigate based on user type and profile completion:
    - Freelancer + profile incomplete: `/onboarding/profile-setup`
    - Freelancer + profile complete: `/dashboard`
    - Client: `/dashboard`
- **Background Events:**
  - users-be publishes `user.logged_in.v1` event
  - Security event logged in `security_events` table
  - Last login timestamp updated

**Error Responses:**

**401 Unauthorized - Invalid Credentials:**
```json
{
  "error": "INVALID_CREDENTIALS",
  "message": "Invalid email or password"
}
```
- Display error toast: "Invalid email or password"
- Clear password field
- Increment failed login attempt counter (backend)
- If 3+ failed attempts:
  - Show CAPTCHA challenge
  - Security event logged
- If 5+ failed attempts:
  - Account locked for 15 minutes
  - Display: "Too many failed attempts. Account locked for 15 minutes."
  - Email sent: "Unusual Login Activity Detected"

**403 Forbidden - Account Locked:**
```json
{
  "error": "ACCOUNT_LOCKED",
  "message": "Account locked due to too many failed login attempts",
  "lockedUntil": "2025-11-07T11:00:00Z"
}
```
- Display error:
  - Title: "Account Locked"
  - Message: "Your account has been locked for security"
  - Countdown: "Unlocked in 14:32"
  - Button: "Reset Password" → `/auth/forgot-password`

**403 Forbidden - Account Suspended:**
```json
{
  "error": "ACCOUNT_SUSPENDED",
  "message": "Your account has been suspended",
  "reason": "Terms violation",
  "suspendedUntil": "2025-11-15T00:00:00Z"
}
```
- Display error:
  - Title: "Account Suspended"
  - Message: "Your account is suspended until {date}"
  - Reason: "{suspension_reason}"
  - Button: "Contact Support" → Opens support email

**403 Forbidden - Account Banned:**
```json
{
  "error": "ACCOUNT_BANNED",
  "message": "Your account has been permanently banned",
  "reason": "Serious terms violation"
}
```
- Display error:
  - Title: "Account Banned"
  - Message: "Your account has been permanently banned"
  - Reason: "{ban_reason}"
  - Button: "Appeal" → `/support/appeal`

**429 Too Many Requests:**
```json
{
  "error": "RATE_LIMIT_EXCEEDED",
  "message": "Too many login attempts. Please try again later.",
  "retryAfter": 900  // seconds
}
```
- Display error toast: "Too many attempts. Try again in 15 minutes."
- Disable form for 15 minutes
- Show countdown timer

**Step 5: OAuth Login Flow (Google)**
- User clicks "Continue with Google"
- Frontend initiates OAuth flow:
  ```
  GET https://keycloak.skillsier.com/realms/skillsier/protocol/openid-connect/auth
  ?client_id=skillsier-web
  &redirect_uri=https://skillsier.com/auth/callback
  &response_type=code
  &scope=openid email profile
  &state={random_state}
  &kc_idp_hint=google
  ```
- Redirect to Google OAuth consent screen
- User authorizes Skillsier
- Google redirects back to Keycloak
- Keycloak exchanges code for Google tokens
- Keycloak retrieves user info from Google
- Keycloak checks if user exists:
  - **Existing user:** Returns Keycloak tokens
  - **New user:** Creates user, returns tokens, triggers registration flow
- Keycloak redirects to: `/auth/callback?code={code}&state={state}`
- Frontend exchanges code for tokens (Step 10 from AUTH-1)
- User logged in and navigated to dashboard

**Step 6: Biometric Login (Mobile)**
- **Prerequisite:** User previously enabled biometric authentication in settings
- User opens mobile app
- App checks if biometric is available and enabled for this user
  ```javascript
  const biometricAvailable = await LocalAuthentication.hasHardwareAsync();
  const biometricEnabled = await SecureStore.getItemAsync('biometric_enabled');
  ```
- If biometric is available and enabled:
  - Show biometric prompt automatically or show "Login with Face ID/Touch ID" button
  - User taps button or biometric prompt appears
  - Native biometric authentication triggered:
    ```javascript
    const result = await LocalAuthentication.authenticateAsync({
      promptMessage: 'Login to Skillsier',
      fallbackLabel: 'Use Password',
    });
    ```
  - If biometric authentication succeeds:
    - Retrieve stored tokens from secure storage:
      ```javascript
      const refreshToken = await SecureStore.getItemAsync('refresh_token');
      ```
    - Call token refresh endpoint to get new access token:
      ```
      POST https://api.skillsier.com/v1/auth/refresh
      Authorization: Bearer {refresh_token}
      ```
    - Store new tokens in secure storage
    - Navigate to dashboard
  - If biometric fails:
    - Show "Use Password" fallback
    - Navigate to login screen with email pre-filled
  - If biometric is cancelled:
    - Show login screen

**Step 7: Device Trust & Remember Me**
- If user checked "Trust this device":
  - Frontend generates device fingerprint:
    ```javascript
    const deviceFingerprint = {
      deviceId: await Application.getIosIdForVendorAsync(), // iOS
      deviceInfo: {
        brand: Device.brand,
        modelName: Device.modelName,
        osName: Device.osName,
        osVersion: Device.osVersion,
      },
      appVersion: Application.nativeApplicationVersion,
    };
    ```
  - Send device fingerprint with login request
  - Backend associates device with user session
  - Device trust stored for 30 days
  - On subsequent logins from trusted device:
    - Skip 2FA if device is trusted
    - Extended session duration (30 days vs 7 days)

**Step 8: Session Refresh**
- Access token expires after 15 minutes
- Frontend automatically refreshes token in background:
  ```javascript
  // Interceptor in axios/fetch
  if (error.response.status === 401 && !isRefreshing) {
    isRefreshing = true;
    const refreshToken = await getRefreshToken();
    const { data } = await axios.post('/v1/auth/refresh', {
      refresh_token: refreshToken
    });
    setAccessToken(data.access_token);
    setRefreshToken(data.refresh_token);
    isRefreshing = false;
    // Retry original request
    return axios(originalRequest);
  }
  ```
- If refresh token is invalid/expired:
  - Clear stored tokens
  - Redirect to login screen
  - Show message: "Your session has expired. Please login again."

**Step 9: Logout**
- User clicks "Logout" in settings or profile menu
- Frontend calls logout endpoint:
  ```
  POST https://api.skillsier.com/v1/auth/logout
  Authorization: Bearer {access_token}
  Body: { 
    refresh_token: "{refresh_token}",
    session_id: "{session_id}"
  }
  ```
- Backend invalidates tokens and session
- Frontend clears all stored tokens and user data:
  ```javascript
  await SecureStore.deleteItemAsync('access_token');
  await SecureStore.deleteItemAsync('refresh_token');
  await SecureStore.deleteItemAsync('user_data');
  await AsyncStorage.clear(); // Clear all cached data
  ```
- Reset app state to initial state
- Navigate to landing/login screen

**Step 10: Session Management**
- Web: Display active sessions in settings:
  - Current session (highlighted)
  - Other sessions with:
    - Device type (Chrome on Windows, Safari on MacBook)
    - Last active time
    - IP address
    - Location (city, country)
    - "Revoke" button for each session
  - "Revoke All Other Sessions" button
- Mobile: Similar display in settings
- User can revoke individual sessions:
  ```
  DELETE https://api.skillsier.com/v1/auth/sessions/{session_id}
  Authorization: Bearer {access_token}
  ```
- User can revoke all other sessions:
  ```
  POST https://api.skillsier.com/v1/auth/sessions/revoke-all
  Authorization: Bearer {access_token}
  Body: { except_current: true }
  ```

#### Branches & Edge Cases

**Account Not Verified:**
- User logs in successfully but email not verified
- Backend returns:
  ```json
  {
    "access_token": "...",
    "refresh_token": "...",
    "user": {
      "id": "...",
      "email_verified": false
    }
  }
  ```
- Frontend navigates to email verification required screen: `/auth/verify-email`
- User prompted to check email or resend verification email
- Limited app functionality until verified (can't post jobs, apply to jobs, etc.)

**2FA Required:**
- After successful email/password authentication
- Backend returns:
  ```json
  {
    "requires_2fa": true,
    "session_id": "temp_session_id",
    "methods": ["AUTHENTICATOR", "SMS"]
  }
  ```
- Frontend navigates to 2FA verification: `/auth/mfa/verify`
- User enters 6-digit code from authenticator app or SMS
- Frontend submits code:
  ```
  POST https://api.skillsier.com/v1/auth/mfa/verify
  Body: { 
    session_id: "temp_session_id",
    code: "123456",
    method: "AUTHENTICATOR",
    trust_device: true
  }
  ```
- On success, receive full access and refresh tokens
- On failure, show error and allow retry (max 3 attempts)
- After 3 failed attempts, lock account for 15 minutes

**Keycloak Integration:**
- All authentication flows go through Keycloak
- Frontend never handles passwords directly
- OAuth2 flow for social login (Google):
  1. Redirect to Keycloak: `/auth/login` → Keycloak login page
  2. Keycloak handles Google OAuth
  3. Keycloak redirects back to: `/auth/callback?code={auth_code}`
  4. Frontend exchanges code for tokens:
     ```
     POST https://keycloak.skillsier.com/realms/skillsier/protocol/openid-connect/token
     Body: {
       grant_type: "authorization_code",
       code: "{auth_code}",
       redirect_uri: "https://skillsier.com/auth/callback",
       client_id: "skillsier-web"
     }
     ```
  5. Receive Keycloak tokens (id_token, access_token, refresh_token)
  6. Exchange Keycloak tokens for Skillsier API tokens:
     ```
     POST https://api.skillsier.com/v1/auth/keycloak/exchange
     Body: { keycloak_access_token: "{token}" }
     ```
  7. Store Skillsier API tokens
  8. Navigate to dashboard

**Username/Password Registration:**
- Users can register with username and password via Keycloak
- Registration form collects:
  - Username (unique)
  - Email
  - Password (with strength requirements)
  - User type (Freelancer/Client)
- Frontend submits to Keycloak registration endpoint
- Keycloak creates user account
- Verification email sent
- User must verify email before full access

#### Notifications

**Login Success (Email):**
- Triggered on every login from new device/location
- Subject: "New Login to Your Skillsier Account"
- Content:
  - Device: Chrome on Windows
  - Location: San Francisco, CA
  - IP: 192.168.1.1
  - Time: Nov 7, 2025 10:30 AM PST
  - "Not you? Secure your account" button → Change password

**Suspicious Login (Email + Push):**
- Triggered when login from unusual location or device
- Subject: "Unusual Login Activity Detected"
- Content:
  - We detected a login from an unusual location
  - Details: {device, location, IP, time}
  - "This was me" or "Secure my account" buttons
- If user clicks "Secure my account":
  - Immediately revoke all sessions
  - Force password reset
  - Send notification to user

**Account Locked (Email):**
- Triggered after 5 failed login attempts
- Subject: "Your Account Has Been Locked"
- Content:
  - Your account was locked due to too many failed login attempts
  - Locked for 30 minutes
  - "Reset Password" button → Password reset flow

**Session Expired (In-App):**
- Show toast/snackbar: "Your session has expired. Please login again."
- Redirect to login screen

#### Analytics

**Events to Track:**
- `auth.login.success` - { method: "email" | "google" | "biometric", device_type, os }
- `auth.login.failed` - { reason: "invalid_credentials" | "account_locked" | "network_error", attempts }
- `auth.logout` - { method: "manual" | "session_expired" }
- `auth.2fa.enabled` - { method: "AUTHENTICATOR" | "SMS" }
- `auth.2fa.verified` - { method, success: boolean, attempts }
- `auth.password.reset.requested` - { method: "email" }
- `auth.password.reset.completed`
- `auth.session.refreshed` - { token_age_seconds }
- `auth.session.revoked` - { revoked_count, revoke_all: boolean }
- `auth.device.trusted` - { device_type, os }
- `auth.biometric.enabled` - { type: "face_id" | "touch_id" }
- `auth.biometric.login.success`
- `auth.biometric.login.failed` - { reason }

**Metrics to Track:**
- Daily Active Users (DAU)
- Monthly Active Users (MAU)
- Login success rate
- 2FA adoption rate
- Biometric login usage (mobile)
- Average session duration
- Session refresh rate
- Account lockout rate
- Password reset rate
- Social login adoption (Google)

#### System Touchpoints

**Backend Services:**
- **Keycloak**: OAuth2 authentication, user management
- **users-be/auth**: Token exchange, session management
- **users-be/security**: 2FA, password management, device trust
- **communications-be/notification**: Login notifications, security alerts

**External Services:**
- **Keycloak**: Identity provider
- **Google OAuth**: Social login
- **SendGrid**: Email notifications
- **FCM/APNS**: Push notifications for security alerts

#### Sources

- combined-fe-folder-strucure.md: `/app/(auth)/login`, `/app/(auth)/mfa`, `/app/(auth)/callback`
- users-be.database-design.md: users table, sessions table, security_events table, two_factor_auth table
- admin-be.database-design.md: admin_users table (for admin authentication)
- users-be.user-stories.md: Authentication domain stories

---

### AUTH-3: Password Reset & Recovery

**ID:** AUTH-3  
**Persona:** Any User  
**Preconditions:** User has registered account with verified email  
**Primary Screens:**
- Web: `/auth/forgot-password`, `/auth/reset-password`
- Mobile: `/auth/forgot-password`, `/auth/reset-password`

#### Flow Steps

**Step 1: Initiate Password Reset**
- User navigates to login screen
- Clicks "Forgot Password?" link
- Navigates to `/auth/forgot-password`
- Screen displays:
  - Email input field
  - "Send Reset Link" button
  - "Remember password?" → Back to login
- User enters email address
- Frontend validates email format
- User submits form
- Frontend calls:
  ```
  POST https://api.skillsier.com/v1/auth/forgot-password
  Body: { 
    email: "user@example.com",
    client_type: "web" | "mobile"
  }
  ```

**Step 2: Backend Processing**
- Backend validates email exists in system
- If email not found:
  - Return success anyway (security: don't reveal if email exists)
  - Don't send email
- If email found:
  - Generate password reset token (UUID)
  - Set expiration time (24 hours from now)
  - Store token in database with user_id association
  - Generate reset URL: `https://skillsier.com/auth/reset-password?token={token}&email={email}`
  - Send email with reset link
  - Return success response

**Step 3: Email Sent Confirmation**
- Frontend receives success response
- Display confirmation screen:
  - "Check Your Email"
  - "We've sent password reset instructions to {email}"
  - "Didn't receive email?" → Resend link
  - Email might take a few minutes to arrive
  - Check spam folder
  - "Back to Login" button
- Allow resend after 60 seconds cooldown:
  - If user clicks "Resend" before 60 seconds, show countdown
  - After 60 seconds, allow resend
  - Max 3 resend attempts per hour per email

**Step 4: User Checks Email**
- User receives email with subject: "Reset Your Skillsier Password"
- Email content:
  - Greeting: "Hi {first_name},"
  - "You requested to reset your password for your Skillsier account."
  - "Click the button below to reset your password:"
  - [Reset Password Button] → Reset URL
  - "Or copy and paste this link: {reset_url}"
  - "This link expires in 24 hours."
  - "If you didn't request this, ignore this email."
  - "Need help? Contact support@skillsier.com"

**Step 5: User Clicks Reset Link**
- User clicks link in email
- Opens browser/app to: `/auth/reset-password?token={token}&email={email}`
- Frontend validates token presence
- Frontend calls verification endpoint:
  ```
  POST https://api.skillsier.com/v1/auth/verify-reset-token
  Body: { 
    token: "{token}",
    email: "{email}"
  }
  ```
- Backend validates:
  - Token exists
  - Token not expired (< 24 hours old)
  - Token not already used
  - Email matches token's user email
- If valid:
  - Return: `{ valid: true }`
  - Frontend shows password reset form
- If invalid:
  - Return: `{ valid: false, reason: "expired" | "used" | "invalid" }`
  - Frontend shows error screen with:
    - "Invalid or Expired Link"
    - Reason message
    - "Request New Reset Link" button → Back to forgot password

**Step 6: Enter New Password**
- User sees password reset form:
  - New Password field (password input with show/hide)
  - Confirm Password field
  - Password strength meter (weak/fair/good/strong)
  - Password requirements checklist:
    - ✓ At least 8 characters
    - ✓ Contains uppercase letter
    - ✓ Contains lowercase letter
    - ✓ Contains number
    - ✓ Contains special character (!@#$%^&*)
  - "Reset Password" button
- User types new password
- Real-time password strength indicator updates
- Requirements checklist updates in real-time
- User confirms password (both fields must match)
- Frontend validates:
  - Passwords match
  - All requirements met
  - New password different from email

**Step 7: Submit New Password**
- User clicks "Reset Password" button
- Frontend calls:
  ```
  POST https://api.skillsier.com/v1/auth/reset-password
  Body: { 
    token: "{token}",
    email: "{email}",
    new_password: "{hashed_password}"
  }
  ```
- Backend:
  - Re-validates token (not expired, not used)
  - Validates password complexity
  - Hashes new password
  - Updates user's password in Keycloak
  - Marks reset token as used
  - Invalidates all existing sessions for this user
  - Creates security event log entry
  - Sends "Password Changed" confirmation email
  - Returns success

**Step 8: Success Confirmation**
- Frontend receives success response
- Display success screen:
  - ✓ Success icon
  - "Password Reset Successful"
  - "Your password has been reset successfully."
  - "All active sessions have been logged out for security."
  - "Continue to Login" button
- Auto-redirect to login after 5 seconds
- User can login immediately with new password

#### Branches & Edge Cases

**Token Expired:**
- User clicks reset link after 24 hours
- Backend returns: `{ valid: false, reason: "expired" }`
- Frontend shows:
  - "This password reset link has expired"
  - "Reset links are valid for 24 hours"
  - "Request New Reset Link" button

**Token Already Used:**
- User clicks reset link that was already used
- Backend returns: `{ valid: false, reason: "used" }`
- Frontend shows:
  - "This reset link has already been used"
  - "Your password may have already been reset"
  - "Try logging in" button → Login screen
  - "Request New Reset Link" button

**Email Not Found (Security):**
- User enters email that doesn't exist
- Backend still returns success (don't reveal if email exists)
- No email is sent
- User sees success screen anyway
- Security: Prevents email enumeration attacks

**Rate Limiting:**
- Max 3 password reset requests per email per hour
- If exceeded:
  ```json
  {
    "error": "RATE_LIMIT_EXCEEDED",
    "message": "Too many password reset requests. Please try again later.",
    "retry_after": 3600
  }
  ```
- Frontend shows error with countdown timer

**Password Too Weak:**
- User submits password that doesn't meet requirements
- Backend returns:
  ```json
  {
    "error": "WEAK_PASSWORD",
    "message": "Password does not meet security requirements",
    "requirements": [...]
  }
  ```
- Frontend shows error with requirements

**Network Errors:**
- If API call fails due to network error
- Show error message: "Unable to connect. Check your internet connection."
- "Retry" button to resubmit

**Mobile Deep Linking:**
- Email reset link opens mobile app if installed
- URL scheme: `skillsier://auth/reset-password?token={token}&email={email}`
- App handles deep link and shows reset password screen
- If app not installed, opens in browser

#### Notifications

**Password Reset Requested (Email):**
- Sent immediately when user requests reset
- Subject: "Reset Your Skillsier Password"
- Contains reset link with 24-hour expiration

**Password Reset Successful (Email):**
- Sent after password successfully reset
- Subject: "Your Skillsier Password Was Changed"
- Content:
  - "Your password was changed on {date} at {time}"
  - Device: {device_info}
  - Location: {location}
  - "If you didn't make this change, contact support immediately"
  - [Secure My Account] button → Contact support

**Suspicious Password Reset (Email + Push):**
- If reset request from unusual location
- Subject: "Unusual Password Reset Request"
- Content:
  - "We detected a password reset request from an unusual location"
  - Details: {location, IP, device}
  - "If this was you, no action needed"
  - "If not, secure your account immediately"

#### Analytics

**Events to Track:**
- `auth.password.reset.requested` - { email, client_type }
- `auth.password.reset.email.sent` - { email }
- `auth.password.reset.link.clicked` - { token, time_since_sent }
- `auth.password.reset.token.verified` - { valid: boolean, reason }
- `auth.password.reset.completed` - { time_since_request }
- `auth.password.reset.failed` - { reason: "expired" | "used" | "weak_password" }
- `auth.password.reset.link.resent` - { resend_count }

**Metrics to Track:**
- Password reset request rate
- Reset completion rate (link clicked → password reset)
- Average time to complete reset
- Token expiration rate
- Failed reset attempts
- Resend rate

#### System Touchpoints

**Backend Services:**
- **users-be/security/recovery**: Password reset logic
- **Keycloak**: Password update
- **communications-be/email**: Send reset emails
- **users-be/security**: Session invalidation

**External Services:**
- **SendGrid**: Email delivery
- **Keycloak**: Identity provider

#### Sources

- combined-fe-folder-strucure.md: `/auth/forgot-password`, `/auth/reset-password`
- users-be.database-design.md: users table, security_events table
- users-be.user-stories.md: Security/recovery domain

---

### AUTH-4: Email Verification

**ID:** AUTH-4  
**Persona:** New User (Freelancer or Client)  
**Preconditions:** User has registered but not verified email  
**Primary Screens:**
- Web: `/auth/verify-email`, `/auth/register/verification`
- Mobile: `/auth/verify-email`

#### Flow Steps

**Step 1: Registration Triggers Verification**
- User completes registration (AUTH-1)
- Backend creates user account
- Generates email verification token
- Stores token in database:
  ```sql
  UPDATE users SET
    email_verification_token = '{token}',
    email_verification_sent_at = NOW(),
    email_verified = FALSE
  WHERE id = '{user_id}'
  ```
- Sends verification email
- Publishes event: `user.registered.v1`

**Step 2: Verification Email Sent**
- Frontend shows post-registration screen:
  - "Check Your Email"
  - "We've sent a verification email to {email}"
  - "Click the link in the email to verify your account"
  - "Didn't receive email?" → Resend button
  - Email might take a few minutes
  - Check spam folder
  - "Continue" button (limited access)

**Step 3: User Checks Email**
- User receives email with subject: "Verify Your Skillsier Email"
- Email content:
  - "Welcome to Skillsier, {first_name}!"
  - "Click the button below to verify your email:"
  - [Verify Email Button] → Verification URL
  - "Or copy and paste: {verification_url}"
  - "This link expires in 7 days"
  - "If you didn't create this account, ignore this email"

**Step 4: User Clicks Verification Link**
- User clicks link: `/auth/register/verification?token={token}`
- Frontend extracts token from URL
- Automatically calls verification endpoint:
  ```
  POST https://api.skillsier.com/v1/auth/verify-email
  Body: { 
    token: "{token}"
  }
  ```
- Show loading spinner during verification

**Step 5: Backend Verifies Email**
- Backend validates token:
  - Token exists
  - Token not expired (< 7 days old)
  - Token not already used
- If valid:
  - Update user record:
    ```sql
    UPDATE users SET
      email_verified = TRUE,
      email_verified_at = NOW(),
      email_verification_token = NULL
    WHERE id = '{user_id}'
    ```
  - Create security event log
  - Publish event: `user.email.verified.v1`
  - Return success with user data
- If invalid:
  - Return error with reason

**Step 6: Success Confirmation**
- Frontend receives success response
- Display success screen:
  - ✓ Success icon
  - "Email Verified!"
  - "Your email has been verified successfully"
  - "Continue to Dashboard" button (if logged in)
  - "Login to Continue" button (if not logged in)
- If user is logged in:
  - Update user state (email_verified = true)
  - Navigate to onboarding or dashboard
- If not logged in:
  - Navigate to login screen with success message

**Step 7: Resend Verification Email**
- If user didn't receive email or link expired
- User clicks "Resend Verification Email" button
- Frontend calls:
  ```
  POST https://api.skillsier.com/v1/auth/resend-verification
  Body: { 
    email: "{email}"
  }
  ```
- Backend:
  - Validates user exists and not already verified
  - Generates new verification token
  - Invalidates old token
  - Sends new verification email
  - Returns success
- Frontend shows:
  - "Verification email sent!"
  - "Check your inbox for a new verification email"
- Cooldown: 60 seconds between resend requests
- Max 5 resend attempts per day

**Step 8: Limited Access Without Verification**
- User can login without verification
- But certain features are restricted:
  - Cannot post jobs (clients)
  - Cannot apply to jobs (freelancers)
  - Cannot send proposals
  - Cannot message users
  - Can view jobs and profiles (read-only)
- Show persistent banner/alert:
  - "Please verify your email to unlock all features"
  - [Verify Email] button → Resend verification email
- Banner appears on all pages until verified

#### Branches & Edge Cases

**Token Expired:**
- User clicks verification link after 7 days
- Backend returns: `{ error: "TOKEN_EXPIRED" }`
- Frontend shows:
  - "Verification Link Expired"
  - "This verification link has expired"
  - "Request New Verification Email" button
  - Button triggers resend flow

**Already Verified:**
- User clicks verification link but email already verified
- Backend returns: `{ error: "ALREADY_VERIFIED" }`
- Frontend shows:
  - "Email Already Verified"
  - "Your email is already verified"
  - "Continue to Dashboard" button (if logged in)
  - "Login" button (if not logged in)

**Invalid Token:**
- User clicks link with invalid/tampered token
- Backend returns: `{ error: "INVALID_TOKEN" }`
- Frontend shows:
  - "Invalid Verification Link"
  - "This verification link is invalid"
  - "Request New Verification Email" button

**Rate Limiting (Resend):**
- User exceeds resend limit (5 per day)
- Backend returns:
  ```json
  {
    "error": "RATE_LIMIT_EXCEEDED",
    "message": "Too many verification emails sent. Try again tomorrow.",
    "retry_after": 86400
  }
  ```
- Frontend shows error with countdown

**Email Not Found (Resend):**
- User enters email that doesn't exist
- Backend returns success anyway (security)
- No email is sent
- Frontend shows success message

**Mobile Deep Linking:**
- Verification link opens mobile app if installed
- URL scheme: `skillsier://auth/verify-email?token={token}`
- App handles deep link and verifies automatically
- If app not installed, opens in browser

**Auto-Login After Verification:**
- If user verifies email while not logged in
- After successful verification, auto-login user:
  - Use token to identify user
  - Generate session tokens
  - Auto-login user
  - Navigate to onboarding/dashboard
- Improves user experience (one less login)

#### Notifications

**Verification Email Sent (Email):**
- Sent immediately after registration
- Subject: "Verify Your Skillsier Email"
- Contains verification link with 7-day expiration

**Email Verified Confirmation (Email):**
- Sent after successful verification
- Subject: "Welcome to Skillsier!"
- Content:
  - "Your email has been verified"
  - "You now have full access to all features"
  - Quick links to:
    - Complete your profile (freelancers)
    - Post your first job (clients)
    - Browse freelancers/jobs

**Reminder Email (Day 3):**
- If user hasn't verified after 3 days
- Subject: "Reminder: Verify Your Skillsier Email"
- Content:
  - "You're almost there! Just verify your email"
  - "Click to verify: {link}"
  - Benefits of verification
  - Link expires in 4 days

**Final Reminder (Day 6):**
- Last reminder before link expires
- Subject: "Last Chance: Verify Your Email (Expires Tomorrow)"
- Urgency: Link expires in 24 hours

#### Analytics

**Events to Track:**
- `auth.email.verification.sent` - { email, user_id }
- `auth.email.verification.link.clicked` - { token, time_since_sent }
- `auth.email.verified` - { user_id, time_since_registration }
- `auth.email.verification.failed` - { reason: "expired" | "invalid" | "already_verified" }
- `auth.email.verification.resent` - { resend_count }
- `auth.email.verification.reminder.sent` - { days_since_registration }

**Metrics to Track:**
- Email verification rate (% of users who verify)
- Average time to verify email
- Verification link click rate
- Resend rate
- Token expiration rate
- Verification funnel drop-off

#### System Touchpoints

**Backend Services:**
- **users-be/user**: User creation, email verification
- **users-be/auth**: Token generation, verification
- **communications-be/email**: Send verification emails

**External Services:**
- **SendGrid**: Email delivery
- **Keycloak**: User email verification status sync

#### Sources

- combined-fe-folder-strucure.md: `/auth/verify-email`, `/auth/register/verification`
- users-be.database-design.md: users table (email_verified, email_verification_token fields)
- users-be.user-stories.md: User domain, verification flow

---

### AUTH-5: Two-Factor Authentication (2FA) Setup & Management

**ID:** AUTH-5  
**Persona:** Any User (Freelancer or Client)  
**Preconditions:** User has active verified account  
**Primary Screens:**
- Web: `/dashboard/settings/security/two-factor/enable`, `/auth/mfa/setup`, `/auth/mfa/verify`
- Mobile: `/settings/security/two-factor`, `/auth/mfa/setup`

#### Flow Steps

**Step 1: Enable 2FA from Settings**
- User navigates to: `/dashboard/settings/security`
- Security settings page displays:
  - Two-Factor Authentication section
  - Current status: "Disabled" (red badge)
  - Description: "Add an extra layer of security to your account"
  - [Enable 2FA] button
  - Benefits list:
    - Protect against unauthorized access
    - Secure your financial transactions
    - Industry-standard security
- User clicks "Enable 2FA" button
- Navigate to: `/auth/mfa/setup`

**Step 2: Choose 2FA Method**
- Setup page displays method selection:
  - **Authenticator App** (Recommended)
    - Use Google Authenticator, Authy, or similar
    - Most secure option
    - Works offline
    - Radio button
  - **SMS** (Available)
    - Receive codes via text message
    - Requires phone number
    - Radio button
  - Currently selected: Authenticator App (default)
- "Continue" button
- User selects preferred method
- Clicks "Continue"

**Step 3a: Setup Authenticator App**
- If user chose Authenticator App:
- Frontend calls:
  ```
  POST https://api.skillsier.com/v1/auth/mfa/setup
  Body: { 
    method: "AUTHENTICATOR",
    user_id: "{user_id}"
  }
  ```
- Backend:
  - Generates TOTP secret key
  - Creates QR code data
  - Returns:
    ```json
    {
      "secret": "JBSWY3DPEHPK3PXP",
      "qr_code": "otpauth://totp/Skillsier:user@email.com?secret=...",
      "backup_codes": ["12345678", "87654321", ...]
    }
    ```
- Frontend displays setup instructions:
  1. **Scan QR Code**
     - QR code image (generated from qr_code data)
     - "Or enter this code manually: {secret}" (with copy button)
  2. **Download Backup Codes**
     - List of 8 backup codes
     - "Save these codes securely"
     - [Download Codes] button → Download as text file
     - [Print Codes] button
     - Warning: "Each code can only be used once"
  3. **Verify Setup**
     - 6-digit code input field
     - "Enter the code from your authenticator app"
     - [Verify] button

**Step 3b: Setup SMS 2FA**
- If user chose SMS:
- Frontend calls:
  ```
  POST https://api.skillsier.com/v1/auth/mfa/setup
  Body: { 
    method: "SMS",
    user_id: "{user_id}"
  }
  ```
- If phone number not on file:
  - Show phone number input
  - Country code selector
  - Phone validation
  - User enters phone number
- Backend:
  - Validates phone number
  - Sends verification SMS with 6-digit code
  - Returns: `{ phone_masked: "+1 *** *** 1234" }`
- Frontend displays:
  - "We sent a verification code to {phone_masked}"
  - 6-digit code input
  - "Didn't receive code?" → Resend (60s cooldown)
  - [Verify] button

**Step 4: Verify 2FA Setup**
- User enters 6-digit code from authenticator app or SMS
- Frontend calls:
  ```
  POST https://api.skillsier.com/v1/auth/mfa/verify-setup
  Body: { 
    user_id: "{user_id}",
    code: "123456",
    method: "AUTHENTICATOR" | "SMS"
  }
  ```
- Backend:
  - Validates code
  - If valid:
    - Enable 2FA for user:
      ```sql
      UPDATE users SET
        two_factor_enabled = TRUE,
        two_factor_method = '{method}',
        two_factor_secret = '{encrypted_secret}'
      WHERE id = '{user_id}'
      ```
    - Store backup codes (encrypted)
    - Create security event
    - Publish event: `user.2fa.enabled.v1`
    - Return success
  - If invalid:
    - Return error with remaining attempts (max 3)

**Step 5: Success Confirmation**
- Display success screen:
  - ✓ Success icon
  - "Two-Factor Authentication Enabled!"
  - "Your account is now more secure"
  - Backup codes reminder:
    - "Don't forget to save your backup codes"
    - [View Backup Codes] button
  - "Done" button → Back to security settings
- Update security settings page:
  - 2FA status: "Enabled" (green badge)
  - Method: "Authenticator App" or "SMS"
  - [Disable 2FA] button
  - [View Backup Codes] button
  - [Change Method] button

**Step 6: Using 2FA on Login**
- User logs in with email/password
- After successful password authentication
- Backend returns:
  ```json
  {
    "requires_2fa": true,
    "session_id": "temp_session_id",
    "method": "AUTHENTICATOR" | "SMS"
  }
  ```
- Frontend navigates to: `/auth/mfa/verify`
- Display 2FA verification screen:
  - "Two-Factor Authentication"
  - "Enter the 6-digit code from your {method}"
  - 6-digit code input (auto-focus, auto-submit on 6 digits)
  - "Use backup code instead" link
  - "Trust this device for 30 days" checkbox
  - [Verify] button
  - "Having trouble?" link → Support

**Step 7: Submit 2FA Code**
- User enters 6-digit code
- Frontend calls:
  ```
  POST https://api.skillsier.com/v1/auth/mfa/verify
  Body: { 
    session_id: "temp_session_id",
    code: "123456",
    trust_device: true|false
  }
  ```
- Backend:
  - Validates TOTP code (30-second time window)
  - If valid:
    - Generate full access + refresh tokens
    - If trust_device=true:
      - Store device fingerprint
      - Trust for 30 days
      - Skip 2FA on this device for 30 days
    - Create security event
    - Return tokens
  - If invalid:
    - Increment failed attempts counter
    - Return error with remaining attempts (max 3)
    - After 3 failures: Lock account for 15 minutes

**Step 8: Login Success**
- Frontend receives tokens
- Store tokens securely
- Navigate to dashboard
- Show success toast: "Login successful"

**Step 9: Using Backup Codes**
- If user clicks "Use backup code instead"
- Display backup code input:
  - "Enter Backup Code"
  - "Enter one of your 8-digit backup codes"
  - 8-digit input field
  - [Verify] button
  - "Lost your backup codes?" → Contact support
- User enters backup code
- Frontend calls:
  ```
  POST https://api.skillsier.com/v1/auth/mfa/verify-backup
  Body: { 
    session_id: "temp_session_id",
    backup_code: "12345678"
  }
  ```
- Backend:
  - Validates backup code exists and not used
  - If valid:
    - Mark code as used (can't be used again)
    - Generate tokens
    - Warn user: "{X} backup codes remaining"
    - If last code used: Prompt to generate new codes
    - Return tokens
  - If invalid:
    - Return error

**Step 10: Disable 2FA**
- User navigates to security settings
- Clicks "Disable 2FA" button
- Display confirmation modal:
  - "Disable Two-Factor Authentication?"
  - Warning: "This will make your account less secure"
  - "Enter your password to confirm" input
  - [Cancel] button
  - [Disable 2FA] button (red)
- User enters password
- Frontend calls:
  ```
  POST https://api.skillsier.com/v1/auth/mfa/disable
  Body: { 
    user_id: "{user_id}",
    password: "{password}"
  }
  ```
- Backend:
  - Validates password
  - If valid:
    - Disable 2FA:
      ```sql
      UPDATE users SET
        two_factor_enabled = FALSE,
        two_factor_method = NULL,
        two_factor_secret = NULL
      WHERE id = '{user_id}'
      ```
    - Delete backup codes
    - Revoke all trusted devices
    - Create security event
    - Send "2FA Disabled" email notification
    - Return success
- Display success message: "2FA has been disabled"
- Update security settings page

#### Branches & Edge Cases

**Lost Authenticator App:**
- User can't access authenticator app
- Options:
  1. Use backup codes (if saved)
  2. Use SMS recovery (if SMS enabled as backup)
  3. Contact support with ID verification

**Lost Backup Codes:**
- User clicks "View Backup Codes" in settings
- Requires password re-authentication
- Display warning: "Anyone with these codes can access your account"
- Show existing codes (with used status)
- [Generate New Codes] button
- Generates new set of 8 codes
- Old codes become invalid

**Change 2FA Method:**
- User clicks "Change Method" in settings
- Requires current 2FA verification
- Then follows setup flow for new method
- Old method is disabled when new method is verified

**Account Lockout (Failed 2FA):**
- After 3 failed 2FA attempts
- Account locked for 15 minutes
- Display: "Account locked due to failed 2FA attempts. Try again in 14:32"
- Countdown timer
- "Use backup code" still available
- "Contact support" link

**Rate Limiting (SMS):**
- Max 5 SMS codes per hour
- If exceeded: Show error and countdown
- Suggest using authenticator app instead

**Device Trust Expiration:**
- After 30 days, device trust expires
- User must verify 2FA again
- No warning, just prompt on next login

**Mobile Biometric + 2FA:**
- If both biometric and 2FA enabled on mobile
- Biometric satisfies 2FA on trusted device
- On new device: Biometric + 2FA both required

#### Notifications

**2FA Enabled (Email):**
- Sent immediately after 2FA setup
- Subject: "Two-Factor Authentication Enabled"
- Content:
  - "2FA was enabled on your account"
  - Method: {method}
  - Time: {timestamp}
  - Device: {device_info}
  - "If this wasn't you, secure your account immediately"

**2FA Disabled (Email):**
- Sent when 2FA disabled
- Subject: "Two-Factor Authentication Disabled"
- Warning: "Your account security has been reduced"
- "Re-enable 2FA" link

**Failed 2FA Attempts (Email):**
- After 2 failed 2FA attempts
- Subject: "Failed 2FA Attempts on Your Account"
- Content:
  - "We detected {count} failed 2FA attempts"
  - Device: {device_info}
  - Location: {location}
  - "If this wasn't you, change your password immediately"

**Account Locked (Email + Push):**
- When account locked due to failed 2FA
- Subject: "Your Account Has Been Locked"
- Content:
  - "Your account was locked for 15 minutes"
  - "Use backup code or contact support"

**Backup Code Used (Email):**
- When backup code is used
- Subject: "Backup Code Used"
- Content:
  - "A backup code was used to login"
  - "{X} backup codes remaining"
  - If < 3 codes: "Consider generating new backup codes"

#### Analytics

**Events to Track:**
- `auth.2fa.setup.started` - { method }
- `auth.2fa.setup.completed` - { method, time_elapsed }
- `auth.2fa.setup.abandoned` - { method, step }
- `auth.2fa.enabled` - { method }
- `auth.2fa.disabled` - { reason }
- `auth.2fa.method.changed` - { old_method, new_method }
- `auth.2fa.verified.success` - { method, device_trusted }
- `auth.2fa.verified.failed` - { method, attempts }
- `auth.2fa.backup_code.used` - { codes_remaining }
- `auth.2fa.backup_codes.regenerated`
- `auth.account.locked.2fa` - { attempts }

**Metrics to Track:**
- 2FA adoption rate (% of users)
- 2FA setup completion rate
- 2FA method distribution (Authenticator vs SMS)
- 2FA verification success rate
- Backup code usage rate
- Account lockout rate due to 2FA
- Device trust adoption rate
- Average time to complete 2FA setup

#### System Touchpoints

**Backend Services:**
- **users-be/security/mfa**: 2FA setup, verification, management
- **users-be/security/two_factor**: TOTP generation, validation
- **communications-be/sms**: SMS code delivery
- **communications-be/email**: Email notifications

**External Services:**
- **Twilio**: SMS delivery
- **SendGrid**: Email notifications
- **TOTP Library**: Time-based OTP generation/validation

#### Sources

- combined-fe-folder-strucure.md: `/auth/mfa/*`, `/settings/security/two-factor/*`
- users-be.database-design.md: users table (two_factor_* fields), two_factor_auth table
- users-be.user-stories.md: Security/2FA domain

---

### AUTH-6: OAuth & Social Login

**ID:** AUTH-6  
**Persona:** Any User (Freelancer or Client)  
**Preconditions:** None (can be new or existing user)  
**Primary Screens:**
- Web: `/auth/login`, `/auth/callback`, `/auth/link-account`
- Mobile: `/auth/login`, Deep link handler

#### Flow Steps

**Step 1: User Chooses Social Login**
- User on login page: `/auth/login`
- Page displays social login options:
  - [Continue with Google] button (Google logo + text)
  - "Or use email/password" separator
  - Email/password form below
- User clicks "Continue with Google" button
- Frontend initiates OAuth flow

**Step 2: Redirect to Keycloak OAuth**
- Frontend redirects to Keycloak OAuth authorization endpoint:
  ```
  GET https://keycloak.skillsier.com/realms/skillsier/protocol/openid-connect/auth
  Parameters:
    client_id: skillsier-web
    redirect_uri: https://skillsier.com/auth/callback
    response_type: code
    scope: openid email profile
    state: {random_state_token}
    kc_idp_hint: google  (hints Keycloak to use Google IDP)
  ```
- State token stored in session for CSRF protection
- User redirected to Keycloak

**Step 3: Keycloak Redirects to Google**
- Keycloak recognizes Google IDP hint
- Redirects to Google OAuth consent screen:
  ```
  GET https://accounts.google.com/o/oauth2/v2/auth
  Parameters:
    client_id: {google_client_id}
    redirect_uri: {keycloak_callback}
    response_type: code
    scope: openid email profile
    state: {state}
  ```
- User sees Google account picker/consent screen

**Step 4: User Authorizes with Google**
- User selects Google account
- If first time: Google shows permission consent screen:
  - "Skillsier wants to access:"
  - ✓ View your email address
  - ✓ View your basic profile info
  - [Cancel] [Allow] buttons
- User clicks "Allow"
- Google redirects back to Keycloak with authorization code

**Step 5: Keycloak Exchanges Code for Google Tokens**
- Keycloak receives callback from Google
- Keycloak exchanges authorization code for tokens:
  ```
  POST https://oauth2.googleapis.com/token
  Body:
    code: {auth_code}
    client_id: {google_client_id}
    client_secret: {google_client_secret}
    redirect_uri: {keycloak_callback}
    grant_type: authorization_code
  ```
- Receives Google access token and ID token

**Step 6: Keycloak Retrieves User Info**
- Keycloak uses Google token to get user profile:
  ```
  GET https://www.googleapis.com/oauth2/v3/userinfo
  Authorization: Bearer {google_access_token}
  ```
- Receives:
  ```json
  {
    "sub": "google_user_id",
    "email": "user@gmail.com",
    "email_verified": true,
    "name": "John Doe",
    "picture": "https://...jpg",
    "given_name": "John",
    "family_name": "Doe"
  }
  ```

**Step 7: Keycloak User Lookup/Creation**
- Keycloak checks if user exists:
  - Search by email: `user@gmail.com`
  - Search by federated identity: `google:{google_user_id}`

**Case A: Existing Keycloak User (Google linked)**
- User found with Google identity
- Keycloak generates authorization code
- Redirects to Skillsier callback:
  ```
  GET https://skillsier.com/auth/callback?code={code}&state={state}
  ```

**Case B: Existing Skillsier User (Email exists, Google not linked)**
- User found by email but no Google identity linked
- Keycloak shows account linking page:
  - "Link Google Account?"
  - "A Skillsier account exists with this email"
  - "Enter your Skillsier password to link accounts"
  - Password input
  - [Cancel] [Link Accounts] buttons
- User enters password
- If password correct:
  - Link Google identity to existing Keycloak user
  - Proceed to authorization
- If password wrong:
  - Show error: "Incorrect password"
  - Retry or cancel

**Case C: New User (No existing account)**
- No user found with email
- Keycloak creates new user:
  ```
  POST /admin/realms/skillsier/users
  Body: {
    email: "user@gmail.com",
    emailVerified: true,  (trust Google's verification)
    firstName: "John",
    lastName: "Doe",
    enabled: true,
    federatedIdentities: [{
      identityProvider: "google",
      userId: "google_user_id",
      userName: "user@gmail.com"
    }]
  }
  ```
- New Keycloak user created
- User needs to complete Skillsier registration

**Step 8: Redirect to Skillsier Callback**
- Keycloak redirects to Skillsier:
  ```
  GET https://skillsier.com/auth/callback?code={auth_code}&state={state}
  ```
- Frontend validates state token (CSRF protection)
- Extracts authorization code

**Step 9: Exchange Code for Skillsier Tokens**
- Frontend exchanges Keycloak code for tokens:
  ```
  POST https://api.skillsier.com/v1/auth/keycloak/exchange
  Body: {
    code: "{auth_code}",
    redirect_uri: "https://skillsier.com/auth/callback",
    state: "{state}"
  }
  ```
- Backend:
  - Validates code with Keycloak
  - Retrieves Keycloak user info
  - Checks if Skillsier user exists (by keycloak_id)

**Case A: Existing Skillsier User**
- User found in Skillsier database
- Generate Skillsier JWT tokens
- Update last_login timestamp
- Create session record
- Return:
  ```json
  {
    "access_token": "...",
    "refresh_token": "...",
    "user": {
      "id": "...",
      "email": "user@gmail.com",
      "first_name": "John",
      "last_name": "Doe",
      "user_type": "FREELANCER",
      "onboarding_completed": true
    }
  }
  ```

**Case B: New Skillsier User**
- User not found in Skillsier database
- Create new Skillsier user:
  ```sql
  INSERT INTO users (
    keycloak_id,
    email,
    first_name,
    last_name,
    email_verified,
    onboarding_completed
  ) VALUES (
    '{keycloak_id}',
    'user@gmail.com',
    'John',
    'Doe',
    TRUE,  -- Trust Google's verification
    FALSE  -- Needs onboarding
  )
  ```
- Generate Skillsier JWT tokens
- Publish event: `user.registered.v1`
- Return tokens + user data (onboarding_completed = false)

**Step 10a: Login Success (Existing User)**
- Frontend receives tokens and user data
- Store tokens securely
- Update user state
- Navigate to dashboard
- Show welcome back message

**Step 10b: Registration Success (New User)**
- Frontend receives tokens and user data
- Notice: `onboarding_completed = false`
- Store tokens
- Navigate to onboarding flow: `/onboarding`
- Complete user type selection and profile setup

**Step 11: Link Additional Social Accounts**
- Existing user wants to add Google login
- Navigate to: `/dashboard/settings/security/linked-accounts`
- Page displays:
  - Connected Accounts section
  - Google: "Not Connected" (red badge)
  - [Connect Google] button
- User clicks "Connect Google"
- Similar OAuth flow but with account linking intent:
  ```
  GET /auth/link?provider=google&return_to=/settings/security
  ```
- After Google auth, link Google identity to existing Keycloak user
- Update Skillsier user record
- Return to settings with success message: "Google account linked"

**Step 12: Unlink Social Account**
- User wants to remove Google login
- Must have password set (can't remove all login methods)
- Display confirmation modal:
  - "Unlink Google Account?"
  - Warning: "You'll need to use your password to login"
  - [Cancel] [Unlink] buttons
- User confirms
- Frontend calls:
  ```
  DELETE https://api.skillsier.com/v1/auth/linked-accounts/google
  ```
- Backend removes federated identity from Keycloak
- Success message: "Google account unlinked"

#### Branches & Edge Cases

**Email Already Exists (Different Provider):**
- User tries to login with Google
- Email exists but linked to Facebook
- Keycloak shows account linking page
- User must verify with password or other linked account
- Then can link new provider

**Email Verification Required:**
- User logs in with Google (email_verified from Google)
- Email is auto-verified in Skillsier
- Skip email verification step

**Failed OAuth (User Cancels):**
- User clicks "Cancel" on Google consent screen
- Google redirects with error:
  ```
  GET /auth/callback?error=access_denied&state={state}
  ```
- Frontend shows message: "Login cancelled. Please try again."
- Return to login page

**Network Errors During OAuth:**
- If Keycloak unreachable
- Show error: "Unable to connect. Please try again."
- Retry button

**State Mismatch (CSRF Attack):**
- State token doesn't match stored state
- Reject authorization
- Show error: "Invalid request. Please try again."
- Clear stored state
- Return to login

**Mobile OAuth Flow:**
- Mobile app initiates OAuth with custom URL scheme
- Redirect URI: `skillsier://auth/callback`
- After OAuth, deep link opens app
- App extracts code and exchanges for tokens
- If app not installed, fallback to web

**Multiple Google Accounts:**
- User has multiple Google accounts in browser
- Google shows account picker
- User selects desired account
- OAuth proceeds with selected account

**Revoked Google Access:**
- User previously granted access but revoked it in Google settings
- On next login attempt, Google prompts for consent again
- User must re-grant permissions

**Expired Keycloak Session:**
- OAuth flow timeout (user took too long)
- Keycloak session expires
- Show error: "Session expired. Please try again."
- Restart OAuth flow

#### Notifications

**New Social Login (Email):**
- Sent when user logs in with new social account
- Subject: "New Login Method Added"
- Content:
  - "Google account was linked to your Skillsier account"
  - Email: {google_email}
  - Time: {timestamp}
  - Device: {device_info}
  - "If this wasn't you, secure your account"

**Social Account Linked (Email):**
- Sent when user links additional social account
- Subject: "Google Account Linked"
- Content:
  - "You can now use your Google account to login"

**Social Account Unlinked (Email):**
- Sent when user unlinks social account
- Subject: "Google Account Unlinked"
- Content:
  - "Google login has been removed"
  - "Use your password to login"

#### Analytics

**Events to Track:**
- `auth.oauth.started` - { provider: "google", client_type }
- `auth.oauth.completed` - { provider, time_elapsed, new_user: boolean }
- `auth.oauth.failed` - { provider, error, step }
- `auth.oauth.cancelled` - { provider, step }
- `auth.account.linked` - { provider, existing_user: true }
- `auth.account.unlinked` - { provider }
- `auth.social.login.success` - { provider, device_type }

**Metrics to Track:**
- Social login adoption rate
- Social login success rate
- Google vs Email/Password login ratio
- Account linking rate
- OAuth funnel drop-off by step
- Average time to complete OAuth flow
- New users via social login (%)

#### System Touchpoints

**Backend Services:**
- **Keycloak**: OAuth2 provider, identity federation
- **users-be/auth**: Token exchange, user creation
- **users-be/user**: User management
- **communications-be/email**: Email notifications

**External Services:**
- **Google OAuth2**: Social authentication
- **Keycloak**: Identity provider and broker

#### Sources

- combined-fe-folder-strucure.md: `/auth/login`, `/auth/callback`, `/settings/security/linked-accounts`
- users-be.database-design.md: users table (keycloak_id field)
- admin-be.database-design.md: admin_users table (keycloak_id field)
- Keycloak: Google Identity Provider configuration

---

### AUTH-7: Account Lockout & Security Events

**ID:** AUTH-7  
**Persona:** Any User  
**Preconditions:** User has active account  
**Primary Screens:**
- Web: `/dashboard/settings/security/events`, `/auth/locked`
- Mobile: `/settings/security/events`, `/auth/locked`

#### Flow Steps

**Step 1: Failed Login Attempts**
- User attempts to login with wrong password
- Backend increments failed_login_attempts counter:
  ```sql
  UPDATE users SET
    failed_login_attempts = failed_login_attempts + 1,
    last_failed_login_at = NOW()
  WHERE email = '{email}'
  ```
- Create security event:
  ```sql
  INSERT INTO security_events (
    user_id, event_type, severity,
    ip_address, user_agent, device_id,
    location_country
  ) VALUES (
    '{user_id}', 'LOGIN_FAILED', 'MEDIUM',
    '{ip}', '{user_agent}', '{device_id}',
    '{country}'
  )
  ```
- Return error with attempts remaining:
  ```json
  {
    "error": "INVALID_CREDENTIALS",
    "message": "Invalid email or password",
    "attempts_remaining": 3
  }
  ```

**Step 2: Progressive Lockout**
- After different thresholds, apply progressive measures:
  
  **After 3 failures:**
  - Show CAPTCHA on login form
  - Require CAPTCHA solve before next attempt
  
  **After 5 failures:**
  - Lock account for 15 minutes
  - Update user record:
    ```sql
    UPDATE users SET
      account_locked_until = NOW() + INTERVAL '15 minutes',
      failed_login_attempts = 5
    WHERE id = '{user_id}'
    ```
  - Create security event: `ACCOUNT_LOCKED`
  - Send email notification
  - Return error:
    ```json
    {
      "error": "ACCOUNT_LOCKED",
      "message": "Too many failed attempts. Account locked for 15 minutes.",
      "locked_until": "2025-11-07T11:00:00Z"
    }
    ```
  
  **After 10 failures (within 24 hours):**
  - Extended lock for 1 hour
  - Require password reset to unlock
  - Send high-priority email + push notification
  
  **After 15 failures:**
  - Manual review required
  - Account stays locked until admin unlocks
  - Security team notified

**Step 3: Locked Account Screen**
- Frontend detects ACCOUNT_LOCKED error
- Navigate to `/auth/locked`
- Display locked account screen:
  - 🔒 Lock icon
  - "Account Temporarily Locked"
  - "Your account has been locked for security reasons"
  - "Too many failed login attempts"
  - Countdown timer: "Unlocked in 14:32"
  - "Reset Password" button → `/auth/forgot-password`
  - "Contact Support" link
- Countdown updates every second
- After countdown expires:
  - Show "Your account is now unlocked"
  - "Try logging in again" button → `/auth/login`

**Step 4: Automatic Unlock**
- Background job runs every minute
- Checks for expired lockouts:
  ```sql
  UPDATE users SET
    account_locked_until = NULL,
    failed_login_attempts = 0
  WHERE account_locked_until < NOW()
    AND account_locked_until IS NOT NULL
  ```
- Create security event: `ACCOUNT_UNLOCKED`
- Send email notification: "Your account has been unlocked"

**Step 5: Suspicious Activity Detection**
- System monitors for suspicious patterns:
  
  **Unusual Location:**
  - User logs in from country different from usual
  - Calculate distance from last login location
  - If > 500 miles in < 1 hour: Flag as suspicious
  - Create event: `SUSPICIOUS_ACTIVITY`
  - Require additional verification (email confirmation)
  
  **Unusual Device:**
  - New device/browser fingerprint
  - Never seen before for this user
  - Create event: `NEW_DEVICE_LOGIN`
  - Send notification email
  - Ask: "Was this you?" with Yes/No buttons
  
  **Impossible Travel:**
  - Login from Location A
  - Then login from Location B 10 minutes later
  - Distance too far to travel in time
  - Create event: `SUSPICIOUS_ACTIVITY` with `CRITICAL` severity
  - Lock account immediately
  - Force password reset
  - Revoke all sessions
  
  **Session Hijacking:**
  - Two simultaneous logins from different IPs/locations
  - Sessions with same token from different IPs
  - Create event: `SESSION_HIJACK_DETECTED`
  - Terminate all sessions immediately
  - Force re-authentication
  - Require 2FA verification

**Step 6: Security Events Dashboard**
- User navigates to: `/dashboard/settings/security/events`
- Display list of recent security events:
  - Event type (icon + label)
  - Timestamp
  - Device info
  - IP address
  - Location (city, country)
  - Status (Success/Failed)
  - Severity badge
- Filters:
  - Event type dropdown (All, Logins, Password Changes, etc.)
  - Date range picker
  - Severity (All, Low, Medium, High, Critical)
- Export button → Download CSV
- Details modal on click:
  - Full event details
  - User agent string
  - Device fingerprint
  - Actions taken
  - "Mark as Safe" button (if suspicious)
  - "Report as Unauthorized" button → Contact support

**Step 7: Event Response Actions**
- **Mark as Safe:**
  - User clicks "Mark as Safe" on suspicious event
  - Frontend calls:
    ```
    POST https://api.skillsier.com/v1/security/events/{event_id}/mark-safe
    ```
  - Backend updates event status
  - Remove any restrictions applied
  - Learn from user feedback (ML model)
  
- **Report as Unauthorized:**
  - User clicks "Report as Unauthorized"
  - Frontend opens support dialog pre-filled with event details
  - User submits report
  - Support team investigates
  - Actions: Lock account, force password reset, revoke sessions

**Step 8: Real-time Security Alerts**
- **Web:** WebSocket connection for real-time alerts
  ```javascript
  const ws = new WebSocket('wss://api.skillsier.com/v1/security/events/stream');
  ws.onmessage = (event) => {
    const securityEvent = JSON.parse(event.data);
    if (securityEvent.severity === 'HIGH' || securityEvent.severity === 'CRITICAL') {
      showSecurityAlert(securityEvent);
    }
  };
  ```
- **Mobile:** Push notification for critical events
- Alert types:
  - Banner notification (persistent until acknowledged)
  - Modal for critical events (blocks interaction)
  - Badge on security settings icon

**Step 9: Automated Response to Threats**
- **Brute Force Detection:**
  - Rate limiting: Max 5 login attempts per IP per minute
  - If exceeded: Block IP for 1 hour
  - WAF rule: Block suspicious IPs automatically
  
- **Credential Stuffing:**
  - Detect multiple accounts accessed from same IP with failed logins
  - Flag accounts for review
  - Require CAPTCHA for all logins from that IP
  
- **Account Takeover Attempt:**
  - Password change from new device/location
  - Require email verification before change takes effect
  - Send email to old address: "Password change requested"
  - 24-hour waiting period for confirmation

#### Branches & Edge Cases

**Account Locked - Password Reset:**
- User clicks "Reset Password" on locked screen
- Immediately unlocks account upon successful password reset
- Skip waiting period
- All sessions revoked

**False Positive (Legitimate Travel):**
- User traveling for work
- Multiple countries in short time
- System flags as suspicious
- User can whitelist locations in settings
- Or mark event as "This was me"

**Manual Unlock by Admin:**
- Support admin can manually unlock account
- Admin dashboard: `/admin/users/{user_id}/security`
- "Unlock Account" button
- Requires admin 2FA confirmation
- Audit log entry created

**IP Whitelist:**
- User can add trusted IPs in settings
- Corporate VPN, home IP, etc.
- Whitelisted IPs skip some security checks
- Still log events but don't trigger alerts

**Account Recovery (Locked > 24 hours):**
- After 24 hours locked, show recovery option
- "Can't access your account?" link
- Identity verification flow:
  - Answer security questions
  - Upload ID document
  - Manual review by support

**Multiple Device Lockout:**
- User locked out on all devices simultaneously
- Sessions terminated across all devices
- Must go through recovery on any device
- Recovery on one device unlocks all

#### Notifications

**Account Locked (Email + Push):**
- Sent immediately when account locked
- Subject: "Your Account Has Been Locked"
- Content:
  - "Too many failed login attempts"
  - "Locked for {duration}"
  - "Reset password to unlock immediately"
  - [Reset Password] button

**Suspicious Login (Email + Push):**
- Sent for logins from unusual location/device
- Subject: "Unusual Login Activity Detected"
- Content:
  - Details: {device, location, IP, time}
  - "Was this you?" with [Yes] [No] buttons
  - "No" → Triggers account lock + password reset

**Account Unlocked (Email):**
- Sent after automatic unlock
- Subject: "Your Account Has Been Unlocked"
- Content:
  - "Your account lockout period has expired"
  - "You can now login"
  - "Secure your account" tips

**Security Event (High/Critical) (Push):**
- Real-time push notification
- "Security Alert: {event_type}"
- Tap to view details

**Weekly Security Summary (Email):**
- Sent every Monday
- Subject: "Your Weekly Security Summary"
- Content:
  - Total logins this week
  - New devices
  - Failed login attempts
  - Security events
  - "View full report" link

#### Analytics

**Events to Track:**
- `security.login.failed` - { attempts, ip, location }
- `security.account.locked` - { reason, duration, attempts }
- `security.account.unlocked` - { method: "automatic" | "password_reset" | "admin" }
- `security.suspicious.detected` - { type, severity, action_taken }
- `security.event.viewed` - { event_id, event_type }
- `security.event.marked_safe` - { event_id }
- `security.event.reported` - { event_id }
- `security.alert.shown` - { severity, type }
- `security.ip.blocked` - { ip, reason }

**Metrics to Track:**
- Account lockout rate
- Average lockout duration
- False positive rate (marked as safe)
- Suspicious activity detection rate
- Time to detect account takeover
- Response time to security events
- Number of blocked IPs
- Brute force attempt rate

#### System Touchpoints

**Backend Services:**
- **users-be/security**: Event tracking, lockout management
- **users-be/security/events**: Security events CRUD
- **admin-be/security**: Admin security management
- **communications-be/notification**: Real-time alerts
- **communications-be/email**: Email notifications

**External Services:**
- **GeoIP**: Location detection
- **MaxMind**: IP intelligence
- **Cloudflare**: WAF, rate limiting, IP blocking
- **SendGrid**: Email delivery
- **FCM/APNS**: Push notifications

#### Sources

- combined-fe-folder-strucure.md: `/settings/security/events`, `/auth/locked`
- users-be.database-design.md: security_events table, sessions table
- users-be.user-stories.md: Security domain, event tracking

---

### AUTH-8: Session Management & Multi-Device

**ID:** AUTH-8  
**Persona:** Any User  
**Preconditions:** User has active account and logged in on one or more devices  
**Primary Screens:**
- Web: `/dashboard/settings/security/sessions`
- Mobile: `/settings/security/sessions`

#### Flow Steps

**Step 1: View Active Sessions**
- User navigates to: `/dashboard/settings/security/sessions`
- Frontend calls:
  ```
  GET https://api.skillsier.com/v1/auth/sessions
  Authorization: Bearer {access_token}
  ```
- Backend queries sessions:
  ```sql
  SELECT * FROM sessions
  WHERE user_id = '{user_id}'
    AND is_active = TRUE
    AND expires_at > NOW()
  ORDER BY last_active_at DESC
  ```
- Returns list of active sessions
- Frontend displays sessions list:
  - **Current Session** (highlighted with badge)
    - Chrome on Windows
    - Last active: Just now
    - Location: San Francisco, CA
    - IP: 192.168.1.1
    - Started: Nov 7, 2025 9:00 AM
  - **Other Sessions:**
    - Safari on iPhone
    - Last active: 5 minutes ago
    - Location: San Francisco, CA
    - IP: 192.168.1.100
    - Started: Nov 6, 2025 3:00 PM
    - [Revoke] button
  - [Revoke All Other Sessions] button at bottom

**Step 2: Session Details**
- Each session displays:
  - Device type + browser (icon + text)
  - Operating system
  - Last active timestamp (relative time)
  - Location (city, country from IP)
  - IP address
  - Session start time
  - Session duration
  - Trusted device status (✓ badge if trusted)
  - [Revoke] button (except current session)
- Click session card to expand details:
  - Full user agent string
  - Device fingerprint
  - Authentication method used (Password, Google, Biometric)
  - 2FA verified status
  - Session ID (masked)
  - Token refresh history
  - Activity log (last 10 activities)

**Step 3: Revoke Single Session**
- User clicks [Revoke] button on specific session
- Confirmation modal:
  - "Revoke Session?"
  - Device: {device_info}
  - "This device will need to login again"
  - [Cancel] [Revoke] buttons
- User confirms
- Frontend calls:
  ```
  DELETE https://api.skillsier.com/v1/auth/sessions/{session_id}
  Authorization: Bearer {access_token}
  ```
- Backend:
  - Validates user owns session
  - Can't revoke current session via this endpoint
  - Updates session:
    ```sql
    UPDATE sessions SET
      is_active = FALSE,
      terminated_at = NOW(),
      logout_reason = 'REVOKED_BY_USER'
    WHERE id = '{session_id}'
      AND user_id = '{user_id}'
    ```
  - Invalidate session tokens in cache:
    ```
    REDIS DEL session:{session_id}
    REDIS DEL user:{user_id}:sessions
    ```
  - Create security event: `SESSION_REVOKED`
  - Send push notification to revoked device
  - Return success
- Frontend:
  - Remove session from list
  - Show success toast: "Session revoked"
  - Update session count

**Step 4: Revoke All Other Sessions**
- User clicks [Revoke All Other Sessions] button
- Confirmation modal:
  - "Revoke All Other Sessions?"
  - "This will sign you out on all devices except this one"
  - "{count} sessions will be revoked"
  - List of sessions to revoke (preview)
  - [Cancel] [Revoke All] buttons
- User confirms
- Frontend calls:
  ```
  POST https://api.skillsier.com/v1/auth/sessions/revoke-all
  Authorization: Bearer {access_token}
  Body: { 
    except_current: true
  }
  ```
- Backend:
  - Get current session ID from access token
  - Deactivate all other sessions:
    ```sql
    UPDATE sessions SET
      is_active = FALSE,
      terminated_at = NOW(),
      logout_reason = 'REVOKED_BY_USER_ALL'
    WHERE user_id = '{user_id}'
      AND id != '{current_session_id}'
      AND is_active = TRUE
    ```
  - Invalidate tokens in cache
  - Create security event: `ALL_SESSIONS_REVOKED`
  - Send push notifications to all revoked devices
  - Return: `{ revoked_count: 5 }`
- Frontend:
  - Update sessions list (only current session remains)
  - Show success message: "5 sessions revoked"

**Step 5: Session Automatic Expiration**
- Backend cron job runs every hour
- Identifies expired sessions:
  ```sql
  SELECT * FROM sessions
  WHERE is_active = TRUE
    AND expires_at < NOW()
  ```
- For each expired session:
  - Update session:
    ```sql
    UPDATE sessions SET
      is_active = FALSE,
      terminated_at = NOW(),
      logout_reason = 'EXPIRED'
    WHERE id = '{session_id}'
    ```
  - Invalidate tokens in cache
  - Create security event: `SESSION_EXPIRED`
- If user tries to use expired access token:
  - API returns 401 Unauthorized
  - Frontend attempts token refresh
  - If refresh token also expired:
    - Clear stored tokens
    - Redirect to login
    - Show message: "Your session has expired. Please login again."

**Step 6: Session Activity Tracking**
- Every API request updates session activity:
  ```sql
  UPDATE sessions SET
    last_active_at = NOW(),
    last_active_ip = '{ip}',
    last_active_user_agent = '{user_agent}'
  WHERE id = '{session_id}'
  ```
- Activity logged:
  - Request path
  - Request method
  - Timestamp
  - IP address
  - Response status
- User can view activity log in session details
- Suspicious activity detection:
  - IP change (session hijacking)
  - User agent change
  - Impossible travel between requests

**Step 7: Multi-Device Sync**
- State synced across devices:
  - When user takes action on Device A (e.g., update profile)
  - Event published: `user.profile.updated.v1`
  - Device B (if logged in) receives push notification via WebSocket/FCM
  - Device B automatically syncs state
- Real-time sync for:
  - Profile updates
  - Settings changes
  - New messages
  - Contract updates
  - Session revocations

**Step 8: Device Trust Management**
- Trusted devices don't require 2FA for 30 days
- User can view trusted devices in settings
- Each trusted device entry shows:
  - Device name/type
  - Trust granted date
  - Expires date
  - [Remove Trust] button
- Remove trust:
  - Device will require 2FA on next login
  - User must re-trust device after 2FA

**Step 9: Session Notifications**
- **New Session Created:**
  - Email sent when new login detected
  - Subject: "New Login to Your Account"
  - Content: Device, location, time
  - "Not you? Secure account" link
  
- **Session Revoked:**
  - Push notification to revoked device
  - "You were logged out by another device"
  - Tap to login again
  
- **All Sessions Revoked:**
  - Push to all revoked devices
  - Email confirmation
  - "You logged out all other devices"

#### Branches & Edge Cases

**Concurrent Session Limit:**
- Platform limit: Max 10 active sessions per user
- If user tries to login when at limit:
  - Option 1: Revoke oldest inactive session automatically
  - Option 2: Show error: "Too many active sessions. Revoke a session to continue."
  - Option 3: Show list, user must manually revoke one

**Session Hijacking Detection:**
- Rapid IP changes for same session
- User agent mismatch
- Impossible travel (different countries in minutes)
- Actions:
  - Immediately terminate session
  - Flag for review
  - Notify user
  - Force re-authentication with 2FA

**Remember Me:**
- If "Remember me" checked on login:
  - Extend session to 30 days
  - Refresh token valid for 30 days
  - Access token still expires after 15 minutes
  - Device marked as trusted
- If not checked:
  - Session expires after 7 days
  - Require re-authentication

**Idle Timeout:**
- After 30 minutes of inactivity:
  - Show "Are you still here?" modal
  - Countdown: 60 seconds
  - User can click "Yes, I'm here" to stay logged in
  - If no response: Auto-logout
- Inactivity detection:
  - No API calls
  - No mouse movement (web)
  - No screen interaction (mobile)

**Lost Device:**
- User reports device lost/stolen
- Contact support
- Support can:
  - View all sessions
  - Identify device by fingerprint
  - Revoke specific session
  - Block device from future logins
  - Force password reset

**Session Transfer (Mobile):**
- User upgrades phone
- Option to transfer session:
  - Old device generates QR code
  - New device scans QR code
  - Session transferred securely
  - Old device session revoked
- Alternative: Manual login on new device

**Session Conflict:**
- User logs in on Device A
- Then logs in on Device B with different user type (client vs freelancer)
- System should handle gracefully:
  - Allow both sessions
  - Or show warning: "You're logged in as {type} on another device"

#### Notifications

**New Session (Email):**
- Sent when new login detected
- Subject: "New Login to Your Skillsier Account"
- Content:
  - Device: Chrome on Windows
  - Location: San Francisco, CA
  - IP: 192.168.1.1
  - Time: Nov 7, 2025 10:30 AM PST
  - "Not you? Secure your account" button → Revoke sessions + change password

**Session Revoked (Push):**
- Sent to revoked device
- Title: "You Were Logged Out"
- Message: "Your session was revoked by another device"
- Tap to view details or login again

**All Sessions Revoked (Email + Push):**
- Email confirmation
- Subject: "All Sessions Logged Out"
- Content:
  - "You logged out all devices"
  - "Login again on any device to continue"

**Suspicious Session Activity (Email + Push):**
- High-priority alert
- Subject: "Suspicious Activity on Your Account"
- Content:
  - "We detected unusual activity on your account"
  - Details: {activity_description}
  - "Secure your account immediately"

#### Analytics

**Events to Track:**
- `session.created` - { device_type, os, browser, location }
- `session.revoked` - { session_id, revoked_by: "user" | "system", reason }
- `session.revoked_all` - { revoked_count }
- `session.expired` - { session_duration, reason }
- `session.activity.tracked` - { path, method, response_time }
- `session.hijack.detected` - { session_id, indicators[] }
- `session.idle_timeout` - { idle_duration }
- `device.trusted` - { device_id, trust_duration }
- `device.trust.removed` - { device_id, reason }

**Metrics to Track:**
- Average active sessions per user
- Session duration (median, p95)
- Device distribution (web vs mobile, browser/OS breakdown)
- Session revocation rate
- Idle timeout rate
- Token refresh frequency
- Session hijacking detection rate
- Concurrent session usage patterns

#### System Touchpoints

**Backend Services:**
- **users-be/security/session**: Session CRUD, revocation
- **users-be/auth**: Token generation, refresh
- **communications-be/notification**: Push notifications
- **communications-be/email**: Email notifications

**External Services:**
- **Redis**: Session storage, token caching
- **GeoIP**: Location detection
- **FCM/APNS**: Push notifications
- **WebSocket**: Real-time session sync

#### Sources

- combined-fe-folder-strucure.md: `/settings/security/sessions`
- users-be.database-design.md: sessions table, security_events table
- users-be.user-stories.md: Security/session domain
