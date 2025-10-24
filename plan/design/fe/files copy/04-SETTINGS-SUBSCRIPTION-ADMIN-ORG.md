# Skillsier Frontend - Complete Folder Structure
## Part 4: Settings, Subscription, Admin Panel & Organization Management

> **CRITICAL**: This document contains ONLY the folder structure, filenames, and inline backend API mappings.  
> **NO CODE IMPLEMENTATIONS** are included per the strict output policy.

---

## Settings Module

```
apps/web/src/app/[locale]/(dashboard)/settings/
│
├── page.tsx                                # Settings overview
│                                            # - Quick access to all settings
│                                            # - Profile completion indicator
│                                            # - Recently changed settings
│                                            # BE: users-be/preferences
│                                            # GET /v1/users/{id}/preferences
│
├── account/
│   ├── page.tsx                            # Account settings
│   │                                        # - Email address
│   │                                        # - Phone number
│   │                                        # - Username
│   │                                        # - Account type
│   │                                        # - Account status
│   │                                        # BE: users-be/user
│   │                                        # GET /v1/users/me
│   │                                        # PATCH /v1/users/{id}/account
│   │
│   ├── email/
│   │   └── change/
│   │       └── page.tsx                    # Change email
│   │                                        # - New email input
│   │                                        # - Password confirmation
│   │                                        # - Verification email sent
│   │                                        # BE: users-be/user
│   │                                        # POST /v1/users/{id}/change-email
│   │
│   ├── phone/
│   │   └── change/
│   │       └── page.tsx                    # Change phone number
│   │                                        # - New phone input
│   │                                        # - OTP verification
│   │                                        # BE: users-be/user
│   │                                        # POST /v1/users/{id}/change-phone
│   │
│   ├── username/
│   │   └── change/
│   │       └── page.tsx                    # Change username
│   │                                        # - New username
│   │                                        # - Availability check
│   │                                        # BE: users-be/user
│   │                                        # POST /v1/users/{id}/change-username
│   │
│   └── delete/
│       └── page.tsx                        # Delete account
│                                            # - Reason for deletion
│                                            # - Data export option
│                                            # - Confirmation (password + checkbox)
│                                            # - GDPR compliance
│                                            # BE: users-be/account
│                                            # POST /v1/users/{id}/delete-account
│                                            # Publishes: AccountDeleted event
│
├── security/
│   ├── page.tsx                            # Security overview
│   │                                        # - Password strength
│   │                                        # - 2FA status
│   │                                        # - Active sessions
│   │                                        # - Login history
│   │                                        # - Security alerts
│   │                                        # BE: users-be/security
│   │                                        # GET /v1/users/{id}/security
│   │
│   ├── password/
│   │   └── change/
│   │       └── page.tsx                    # Change password
│   │                                        # - Current password
│   │                                        # - New password
│   │                                        # - Password strength meter
│   │                                        # BE: users-be/security
│   │                                        # POST /v1/users/{id}/change-password
│   │
│   ├── two-factor/
│   │   ├── page.tsx                        # 2FA settings
│   │   │                                    # - Enable/disable 2FA
│   │   │                                    # - Backup codes
│   │   │                                    # - Trusted devices
│   │   │                                    # BE: users-be/security/mfa
│   │   │                                    # GET /v1/users/{id}/mfa
│   │   │
│   │   ├── enable/
│   │   │   └── page.tsx                    # Enable 2FA
│   │   │                                    # - QR code scan
│   │   │                                    # - Verify setup
│   │   │                                    # - Save backup codes
│   │   │                                    # BE: users-be/security/mfa
│   │   │                                    # POST /v1/users/{id}/mfa/enable
│   │   │
│   │   └── disable/
│   │       └── page.tsx                    # Disable 2FA
│   │                                        # - Password confirmation
│   │                                        # - 2FA code verification
│   │                                        # BE: users-be/security/mfa
│   │                                        # POST /v1/users/{id}/mfa/disable
│   │
│   ├── sessions/
│   │   ├── page.tsx                        # Active sessions
│   │   │                                    # - List all active sessions
│   │   │                                    # - Device info
│   │   │                                    # - Location
│   │   │                                    # - Last active
│   │   │                                    # - Revoke session
│   │   │                                    # BE: users-be/security/session
│   │   │                                    # GET /v1/users/{id}/sessions
│   │   │                                    # DELETE /v1/users/{id}/sessions/{session_id}
│   │   │
│   │   └── revoke-all/
│   │       └── page.tsx                    # Revoke all sessions (except current)
│   │                                        # BE: users-be/security/session
│   │                                        # POST /v1/users/{id}/sessions/revoke-all
│   │
│   └── login-history/
│       └── page.tsx                        # Login history
│                                            # - Recent logins
│                                            # - Failed attempts
│                                            # - Device/location info
│                                            # BE: users-be/security/audit
│                                            # GET /v1/users/{id}/login-history
│
├── privacy/
│   ├── page.tsx                            # Privacy settings
│   │                                        # - Profile visibility
│   │                                        # - Search visibility
│   │                                        # - Data sharing preferences
│   │                                        # - Contact preferences
│   │                                        # BE: users-be/privacy
│   │                                        # GET /v1/users/{id}/privacy
│   │                                        # PATCH /v1/users/{id}/privacy
│   │
│   ├── blocked-users/
│   │   ├── page.tsx                        # Blocked users list
│   │   │                                    # - List blocked users
│   │   │                                    # - Unblock option
│   │   │                                    # BE: users-be/privacy
│   │   │                                    # GET /v1/users/{id}/blocked-users
│   │   │                                    # DELETE /v1/users/{id}/blocked-users/{user_id}
│   │   │
│   │   └── add/
│   │       └── page.tsx                    # Block user
│   │                                        # BE: users-be/privacy
│   │                                        # POST /v1/users/{id}/blocked-users
│   │
│   └── data-export/
│       ├── page.tsx                        # Data export (GDPR)
│       │                                    # - Request data export
│       │                                    # - Export history
│       │                                    # - Download exports
│       │                                    # BE: users-be/privacy
│       │                                    # POST /v1/users/{id}/data-export
│       │                                    # GET /v1/users/{id}/data-exports
│       │
│       └── request/
│           └── page.tsx                    # Request new data export
│                                            # - Select data categories
│                                            # - Format (JSON, CSV, PDF)
│                                            # BE: users-be/privacy
│                                            # POST /v1/users/{id}/data-export/request
│
├── notifications/
│   ├── page.tsx                            # Notification preferences
│   │                                        # - Email notifications
│   │                                        # - Push notifications (web/mobile)
│   │                                        # - SMS notifications
│   │                                        # - In-app notifications
│   │                                        # - Notification frequency
│   │                                        # - Quiet hours
│   │                                        # BE: communications-be/preferences
│   │                                        # GET /v1/notifications/preferences
│   │                                        # PUT /v1/notifications/preferences
│   │
│   ├── email/
│   │   └── page.tsx                        # Email notification settings
│   │                                        # - Per-category toggles
│   │                                        # - Digest preferences
│   │                                        # BE: communications-be/preferences
│   │                                        # PUT /v1/notifications/email-preferences
│   │
│   ├── push/
│   │   └── page.tsx                        # Push notification settings
│   │                                        # - Device management
│   │                                        # - Per-category toggles
│   │                                        # BE: communications-be/preferences
│   │                                        # PUT /v1/notifications/push-preferences
│   │
│   └── digest/
│       └── page.tsx                        # Email digest settings
│                                            # - Daily/weekly/monthly
│                                            # - Content preferences
│                                            # BE: communications-be/preferences
│                                            # PUT /v1/notifications/digest-preferences
│
├── preferences/
│   ├── page.tsx                            # General preferences
│   │                                        # - Language
│   │                                        # - Timezone
│   │                                        # - Date format
│   │                                        # - Currency
│   │                                        # - Theme (light/dark/system)
│   │                                        # BE: users-be/preferences
│   │                                        # GET /v1/users/{id}/preferences
│   │                                        # PATCH /v1/users/{id}/preferences
│   │
│   ├── language/
│   │   └── page.tsx                        # Language settings
│   │                                        # - Interface language
│   │                                        # - Content languages
│   │                                        # BE: users-be/preferences
│   │                                        # PATCH /v1/users/{id}/preferences/language
│   │
│   └── accessibility/
│       └── page.tsx                        # Accessibility settings
│                                            # - Screen reader optimizations
│                                            # - High contrast mode
│                                            # - Font size
│                                            # - Keyboard shortcuts
│                                            # - Motion preferences
│                                            # BE: users-be/preferences
│                                            # PATCH /v1/users/{id}/preferences/accessibility
│
├── integrations/
│   ├── page.tsx                            # Third-party integrations
│   │                                        # - Connected accounts
│   │                                        # - API access
│   │                                        # - Webhooks
│   │                                        # BE: users-be/integrations
│   │                                        # GET /v1/users/{id}/integrations
│   │
│   ├── calendar/
│   │   └── page.tsx                        # Calendar integration
│   │                                        # - Google Calendar
│   │                                        # - Outlook Calendar
│   │                                        # - Sync settings
│   │                                        # BE: users-be/integrations
│   │                                        # POST /v1/users/{id}/integrations/calendar
│   │
│   ├── slack/
│   │   └── page.tsx                        # Slack integration
│   │                                        # - Connect Slack workspace
│   │                                        # - Notification channels
│   │                                        # BE: users-be/integrations
│   │                                        # POST /v1/users/{id}/integrations/slack
│   │
│   └── webhooks/
│       ├── page.tsx                        # Webhook management
│       │                                    # - List webhooks
│       │                                    # - Create webhook
│       │                                    # - Event selection
│       │                                    # BE: users-be/integrations
│       │                                    # GET /v1/users/{id}/webhooks
│       │
│       ├── create/
│       │   └── page.tsx                    # Create webhook
│       │                                    # - Webhook URL
│       │                                    # - Secret key
│       │                                    # - Events to subscribe
│       │                                    # BE: users-be/integrations
│       │                                    # POST /v1/users/{id}/webhooks
│       │
│       └── [webhookId]/
│           └── page.tsx                    # Webhook detail
│                                            # - Delivery logs
│                                            # - Test webhook
│                                            # - Edit/delete
│                                            # BE: users-be/integrations
│                                            # GET /v1/users/{id}/webhooks/{webhook_id}
│                                            # PUT /v1/users/{id}/webhooks/{webhook_id}
│                                            # DELETE /v1/users/{id}/webhooks/{webhook_id}
│
└── developer/
    ├── page.tsx                            # Developer settings
    │                                        # - API keys
    │                                        # - OAuth applications
    │                                        # - API usage stats
    │                                        # BE: users-be/developer
    │                                        # GET /v1/users/{id}/developer
    │
    ├── api-keys/
    │   ├── page.tsx                        # API keys list
    │   │                                    # - Active keys
    │   │                                    # - Create new key
    │   │                                    # - Revoke key
    │   │                                    # BE: users-be/developer
    │   │                                    # GET /v1/users/{id}/api-keys
    │   │
    │   └── create/
    │       └── page.tsx                    # Create API key
    │                                        # - Key name
    │                                        # - Permissions/scopes
    │                                        # - Expiration
    │                                        # BE: users-be/developer
    │                                        # POST /v1/users/{id}/api-keys
    │
    └── oauth-apps/
        ├── page.tsx                        # OAuth applications
        │                                    # - Registered apps
        │                                    # - Create new app
        │                                    # BE: users-be/developer
        │                                    # GET /v1/users/{id}/oauth-apps
        │
        ├── create/
        │   └── page.tsx                    # Create OAuth app
        │                                    # - App name
        │                                    # - Redirect URIs
        │                                    # - Scopes
        │                                    # BE: users-be/developer
        │                                    # POST /v1/users/{id}/oauth-apps
        │
        └── [appId]/
            └── page.tsx                    # OAuth app detail
                                            # - Client ID/secret
                                            # - Edit/delete
                                            # - Usage stats
                                            # BE: users-be/developer
                                            # GET /v1/users/{id}/oauth-apps/{app_id}
```

---

## Subscription Management Module

```
apps/web/src/app/[locale]/(dashboard)/subscription/
│
├── page.tsx                                # Subscription overview
│                                            # - Current plan details
│                                            # - Usage stats
│                                            # - Next billing date
│                                            # - Upgrade/downgrade options
│                                            # - Connects balance
│                                            # BE: subscriptions-be/subscriptions
│                                            # GET /v1/subscriptions/current
│                                            # BE: subscriptions-be/entitlements
│                                            # GET /v1/entitlements
│                                            # BE: subscriptions-be/connects
│                                            # GET /v1/connects/balance
│
├── plans/
│   ├── page.tsx                            # Available plans
│   │                                        # - Plan comparison table
│   │                                        # - Feature matrix
│   │                                        # - Pricing tiers
│   │                                        # - Billing periods (monthly, annual)
│   │                                        # - Free trial info
│   │                                        # BE: subscriptions-be/plans
│   │                                        # GET /v1/plans
│   │
│   └── compare/
│       └── page.tsx                        # Plan comparison
│                                            # - Side-by-side comparison
│                                            # - Feature highlights
│                                            # BE: subscriptions-be/plans
│                                            # GET /v1/plans/compare
│
├── upgrade/
│   ├── page.tsx                            # Upgrade plan
│   │                                        # - Select new plan
│   │                                        # - Billing period
│   │                                        # - Proration calculation
│   │                                        # - Payment method
│   │                                        # - Confirm upgrade
│   │                                        # BE: subscriptions-be/subscriptions
│   │                                        # POST /v1/subscriptions/upgrade
│   │                                        # Body: { plan_id, billing_period }
│   │                                        # Publishes: SubscriptionUpgraded event
│   │
│   └── confirm/
│       └── page.tsx                        # Confirm upgrade
│                                            # - Upgrade summary
│                                            # - Payment processing
│                                            # BE: subscriptions-be/subscriptions
│                                            # POST /v1/subscriptions/{sub_id}/confirm-upgrade
│
├── downgrade/
│   └── page.tsx                            # Downgrade plan
│                                            # - Select new plan
│                                            # - Effective date (end of billing period)
│                                            # - Feature comparison
│                                            # - Confirm downgrade
│                                            # BE: subscriptions-be/subscriptions
│                                            # POST /v1/subscriptions/downgrade
│                                            # Publishes: SubscriptionDowngraded event
│
├── cancel/
│   └── page.tsx                            # Cancel subscription
│                                            # - Cancellation reason
│                                            # - Feedback
│                                            # - Immediate vs. end of period
│                                            # - Refund eligibility
│                                            # - Data retention info
│                                            # BE: subscriptions-be/subscriptions
│                                            # POST /v1/subscriptions/{sub_id}/cancel
│                                            # Publishes: SubscriptionCancelled event
│
├── reactivate/
│   └── page.tsx                            # Reactivate subscription
│                                            # - Select plan
│                                            # - Payment method
│                                            # BE: subscriptions-be/subscriptions
│                                            # POST /v1/subscriptions/reactivate
│
├── connects/
│   ├── page.tsx                            # Connects overview
│   │                                        # - Current balance
│   │                                        # - Usage history
│   │                                        # - Included in plan
│   │                                        # - Purchase more
│   │                                        # BE: subscriptions-be/connects
│   │                                        # GET /v1/connects/balance
│   │                                        # GET /v1/connects/usage
│   │
│   ├── purchase/
│   │   └── page.tsx                        # Purchase connects
│   │                                        # - Select package
│   │                                        # - Pricing options
│   │                                        # - Bulk discounts
│   │                                        # - Payment method
│   │                                        # BE: subscriptions-be/connects
│   │                                        # POST /v1/connects/purchase
│   │                                        # Body: { package_id, quantity }
│   │                                        # Publishes: ConnectsPurchased event
│   │
│   └── history/
│       └── page.tsx                        # Connects usage history
│                                            # - Connects spent (proposals)
│                                            # - Connects added (purchases/plan)
│                                            # - Balance over time
│                                            # BE: subscriptions-be/connects
│                                            # GET /v1/connects/transactions
│
├── addons/
│   ├── page.tsx                            # Available addons
│   │                                        # - List of addons
│   │                                        # - Active addons
│   │                                        # - Purchase addon
│   │                                        # BE: subscriptions-be/addons
│   │                                        # GET /v1/addons
│   │                                        # GET /v1/subscriptions/{sub_id}/addons
│   │
│   └── [addonId]/
│       ├── purchase/
│       │   └── page.tsx                    # Purchase addon
│       │                                    # BE: subscriptions-be/addons
│       │                                    # POST /v1/subscriptions/{sub_id}/addons
│       │
│       └── cancel/
│           └── page.tsx                    # Cancel addon
│                                            # BE: subscriptions-be/addons
│                                            # DELETE /v1/subscriptions/{sub_id}/addons/{addon_id}
│
├── billing-history/
│   ├── page.tsx                            # Billing history
│   │                                        # - Past invoices
│   │                                        # - Payment history
│   │                                        # - Download invoices
│   │                                        # BE: subscriptions-be/invoices
│   │                                        # GET /v1/subscriptions/{sub_id}/invoices
│   │
│   └── [invoiceId]/
│       └── page.tsx                        # Invoice detail
│                                            # - Invoice details
│                                            # - Download PDF
│                                            # BE: subscriptions-be/invoices
│                                            # GET /v1/invoices/{invoice_id}
│                                            # GET /v1/invoices/{invoice_id}/pdf
│
├── usage/
│   └── page.tsx                            # Usage statistics
│                                            # - Jobs posted
│                                            # - Proposals submitted
│                                            # - Storage used
│                                            # - API calls
│                                            # - Usage vs. limits
│                                            # BE: subscriptions-be/usage
│                                            # GET /v1/subscriptions/usage
│
└── trial/
    ├── page.tsx                            # Trial status
    │                                        # - Trial end date
    │                                        # - Days remaining
    │                                        # - Trial features
    │                                        # - Upgrade prompt
    │                                        # BE: subscriptions-be/trials
    │                                        # GET /v1/trials/current
    │
    └── convert/
        └── page.tsx                        # Convert trial to paid
                                            # - Select plan
                                            # - Payment method
                                            # - Apply promotion code
                                            # BE: subscriptions-be/trials
                                            # POST /v1/trials/{trial_id}/convert
```

---

## Organization Management Module (for Clients)

```
apps/web/src/app/[locale]/(dashboard)/organization/
│
├── page.tsx                                # Organization overview
│                                            # - Company details
│                                            # - Team members
│                                            # - Spending overview
│                                            # - Active contracts
│                                            # BE: users-be/organization
│                                            # GET /v1/organizations/{org_id}
│
├── settings/
│   ├── page.tsx                            # Organization settings
│   │                                        # - Company name
│   │                                        # - Industry
│   │                                        # - Company size
│   │                                        # - Website
│   │                                        # - Logo
│   │                                        # BE: users-be/organization
│   │                                        # PATCH /v1/organizations/{org_id}
│   │
│   └── billing/
│       └── page.tsx                        # Organization billing
│                                            # - Billing profile
│                                            # - Tax information
│                                            # - Payment methods
│                                            # BE: financial-be/billing_profile
│                                            # GET /v1/organizations/{org_id}/billing-profile
│
├── team/
│   ├── page.tsx                            # Team members list
│   │                                        # - Active members
│   │                                        # - Pending invitations
│   │                                        # - Roles
│   │                                        # BE: users-be/team
│   │                                        # GET /v1/organizations/{org_id}/members
│   │
│   ├── invite/
│   │   └── page.tsx                        # Invite team member
│   │                                        # - Email address
│   │                                        # - Role selection
│   │                                        # - Permissions
│   │                                        # BE: users-be/team
│   │                                        # POST /v1/organizations/{org_id}/members/invite
│   │                                        # BE: communications-be
│   │                                        # Sends invitation email
│   │
│   ├── [memberId]/
│   │   ├── page.tsx                        # Member detail
│   │   │                                    # - Member info
│   │   │                                    # - Role & permissions
│   │   │                                    # - Activity
│   │   │                                    # BE: users-be/team
│   │   │                                    # GET /v1/organizations/{org_id}/members/{member_id}
│   │   │
│   │   ├── edit/
│   │   │   └── page.tsx                    # Edit member
│   │   │                                    # - Change role
│   │   │                                    # - Update permissions
│   │   │                                    # BE: users-be/team
│   │   │                                    # PATCH /v1/organizations/{org_id}/members/{member_id}
│   │   │
│   │   └── remove/
│   │       └── page.tsx                    # Remove member
│   │                                        # BE: users-be/team
│   │                                        # DELETE /v1/organizations/{org_id}/members/{member_id}
│   │
│   └── roles/
│       ├── page.tsx                        # Roles management
│       │                                    # - List custom roles
│       │                                    # - Create role
│       │                                    # BE: users-be/role
│       │                                    # GET /v1/organizations/{org_id}/roles
│       │
│       ├── create/
│       │   └── page.tsx                    # Create custom role
│       │                                    # - Role name
│       │                                    # - Permissions selection
│       │                                    # BE: users-be/role
│       │                                    # POST /v1/organizations/{org_id}/roles
│       │
│       └── [roleId]/
│           └── edit/
│               └── page.tsx                # Edit role
│                                            # BE: users-be/role
│                                            # PUT /v1/organizations/{org_id}/roles/{role_id}
│                                            # DELETE /v1/organizations/{org_id}/roles/{role_id}
│
├── spending/
│   ├── page.tsx                            # Spending overview
│   │                                        # - Total spending
│   │                                        # - By project
│   │                                        # - By freelancer
│   │                                        # - By time period
│   │                                        # BE: financial-be/reports
│   │                                        # GET /v1/organizations/{org_id}/spending
│   │
│   └── budgets/
│       ├── page.tsx                        # Budget management
│       │                                    # - Set budgets
│       │                                    # - Budget alerts
│       │                                    # BE: financial-be/budget
│       │                                    # GET /v1/organizations/{org_id}/budgets
│       │
│       └── create/
│           └── page.tsx                    # Create budget
│                                            # BE: financial-be/budget
│                                            # POST /v1/organizations/{org_id}/budgets
│
└── analytics/
    └── page.tsx                            # Organization analytics
                                            # - Hiring metrics
                                            # - Freelancer performance
                                            # - Cost per hire
                                            # - Time to hire
                                            # BE: analytics-be
                                            # GET /v1/analytics/organization/{org_id}
```

---

## Admin Panel Module

```
apps/web/src/app/[locale]/(dashboard)/admin/
│
├── layout.tsx                              # Admin layout (RBAC guard)
│                                            # - Only ADMIN, SUPER_ADMIN, MODERATOR
│                                            # - Admin navigation sidebar
│
├── page.tsx                                # Admin dashboard
│                                            # - Key metrics
│                                            # - Pending moderation queue
│                                            # - Recent admin actions
│                                            # - System alerts
│                                            # BE: admin-be/dashboard
│                                            # GET /v1/admin/dashboard
│
├── users/
│   ├── page.tsx                            # Users management
│   │                                        # - User list
│   │                                        # - Search users
│   │                                        # - Filter by status/type
│   │                                        # - Bulk actions
│   │                                        # BE: admin-be/users
│   │                                        # GET /v1/admin/users?filters={...}
│   │
│   ├── [userId]/
│   │   ├── page.tsx                        # User detail (admin view)
│   │   │                                    # - Full user info
│   │   │                                    # - Activity logs
│   │   │                                    # - Violations
│   │   │                                    # - Admin actions
│   │   │                                    # BE: admin-be/users
│   │   │                                    # GET /v1/admin/users/{user_id}
│   │   │
│   │   ├── suspend/
│   │   │   └── page.tsx                    # Suspend user
│   │   │                                    # - Suspension reason
│   │   │                                    # - Duration
│   │   │                                    # - Notify user
│   │   │                                    # BE: admin-be/users
│   │   │                                    # POST /v1/admin/users/{user_id}/suspend
│   │   │                                    # Publishes: UserSuspended event
│   │   │
│   │   ├── ban/
│   │   │   └── page.tsx                    # Ban user
│   │   │                                    # - Ban reason
│   │   │                                    # - Permanent/temporary
│   │   │                                    # - Related accounts
│   │   │                                    # BE: admin-be/users
│   │   │                                    # POST /v1/admin/users/{user_id}/ban
│   │   │                                    # Publishes: UserBanned event
│   │   │
│   │   ├── warn/
│   │   │   └── page.tsx                    # Warn user
│   │   │                                    # BE: admin-be/users
│   │   │                                    # POST /v1/admin/users/{user_id}/warn
│   │   │
│   │   └── verify/
│   │       └── page.tsx                    # Verify user (manual)
│   │                                        # BE: admin-be/users
│   │                                        # POST /v1/admin/users/{user_id}/verify
│   │
│   └── bulk-actions/
│       └── page.tsx                        # Bulk user actions
│                                            # BE: admin-be/users
│                                            # POST /v1/admin/users/bulk-action
│
├── moderation/
│   ├── page.tsx                            # Moderation queue
│   │                                        # - Flagged content
│   │                                        # - Reports
│   │                                        # - Priority queue
│   │                                        # BE: admin-be/moderation
│   │                                        # GET /v1/admin/moderation/queue
│   │
│   ├── jobs/
│   │   ├── page.tsx                        # Flagged jobs
│   │   │                                    # BE: admin-be/moderation
│   │   │                                    # GET /v1/admin/moderation/jobs
│   │   │
│   │   └── [jobId]/
│   │       └── review/
│   │           └── page.tsx                # Review flagged job
│   │                                        # - Job content
│   │                                        # - Flag reason
│   │                                        # - Actions: Approve/Remove/Hide
│   │                                        # BE: admin-be/moderation
│   │                                        # POST /v1/admin/moderation/jobs/{job_id}/review
│   │
│   ├── proposals/
│   │   └── [proposalId]/
│   │       └── review/
│   │           └── page.tsx                # Review flagged proposal
│   │                                        # BE: admin-be/moderation
│   │                                        # POST /v1/admin/moderation/proposals/{proposal_id}/review
│   │
│   ├── reviews/
│   │   └── [reviewId]/
│   │       └── review/
│   │           └── page.tsx                # Review flagged review
│   │                                        # BE: admin-be/moderation
│   │                                        # POST /v1/admin/moderation/reviews/{review_id}/review
│   │
│   └── messages/
│       └── [messageId]/
│           └── review/
│               └── page.tsx                # Review flagged message
│                                            # BE: admin-be/moderation
│                                            # POST /v1/admin/moderation/messages/{message_id}/review
│
├── kyc/
│   ├── page.tsx                            # KYC cases queue
│   │                                        # - Pending cases
│   │                                        # - Approved/rejected cases
│   │                                        # BE: admin-be/kyc_case
│   │                                        # GET /v1/admin/kyc/cases
│   │
│   └── [caseId]/
│       ├── page.tsx                        # KYC case detail
│       │                                    # - Submitted documents
│       │                                    # - Document viewer
│       │                                    # - User information
│       │                                    # - Case history
│       │                                    # BE: admin-be/kyc_case
│       │                                    # GET /v1/admin/kyc/cases/{case_id}
│       │
│       ├── approve/
│       │   └── page.tsx                    # Approve KYC
│       │                                    # BE: admin-be/kyc_case
│       │                                    # POST /v1/admin/kyc/cases/{case_id}/approve
│       │                                    # Publishes: KYCApproved event
│       │
│       └── reject/
│           └── page.tsx                    # Reject KYC
│                                            # - Rejection reason
│                                            # - Required actions
│                                            # BE: admin-be/kyc_case
│                                            # POST /v1/admin/kyc/cases/{case_id}/reject
│
├── business-verification/
│   ├── page.tsx                            # Business verification queue
│   │                                        # BE: admin-be/business_verification
│   │                                        # GET /v1/admin/business-verification/cases
│   │
│   └── [caseId]/
│       ├── page.tsx                        # Business verification case
│       │                                    # BE: admin-be/business_verification
│       │                                    # GET /v1/admin/business-verification/{case_id}
│       │
│       └── review/
│           └── page.tsx                    # Review business documents
│                                            # BE: admin-be/business_verification
│                                            # POST /v1/admin/business-verification/{case_id}/review
│
├── disputes/
│   ├── page.tsx                            # Disputes management
│   │                                        # - Open disputes
│   │                                        # - Assigned to me
│   │                                        # - Resolved disputes
│   │                                        # BE: admin-be/disputes
│   │                                        # GET /v1/admin/disputes
│   │
│   └── [disputeId]/
│       ├── page.tsx                        # Dispute detail (admin)
│       │                                    # - Dispute information
│       │                                    # - Evidence from both parties
│       │                                    # - Communication history
│       │                                    # - Admin notes
│       │                                    # BE: admin-be/disputes
│       │                                    # GET /v1/admin/disputes/{dispute_id}
│       │
│       ├── assign/
│       │   └── page.tsx                    # Assign dispute to admin
│       │                                    # BE: admin-be/disputes
│       │                                    # POST /v1/admin/disputes/{dispute_id}/assign
│       │
│       ├── resolve/
│       │   └── page.tsx                    # Resolve dispute
│       │                                    # - Resolution decision
│       │                                    # - Financial settlement
│       │                                    # - Explanation
│       │                                    # BE: admin-be/disputes
│       │                                    # POST /v1/admin/disputes/{dispute_id}/resolve
│       │                                    # Publishes: DisputeResolved event
│       │
│       └── escalate/
│           └── page.tsx                    # Escalate dispute
│                                            # BE: admin-be/disputes
│                                            # POST /v1/admin/disputes/{dispute_id}/escalate
│
├── refunds/
│   ├── page.tsx                            # Refund cases
│   │                                        # - Pending refund requests
│   │                                        # - Processed refunds
│   │                                        # BE: admin-be/refund_case
│   │                                        # GET /v1/admin/refunds
│   │
│   └── [caseId]/
│       ├── page.tsx                        # Refund case detail
│       │                                    # BE: admin-be/refund_case
│       │                                    # GET /v1/admin/refunds/{case_id}
│       │
│       ├── approve/
│       │   └── page.tsx                    # Approve refund
│       │                                    # BE: admin-be/refund_case
│       │                                    # POST /v1/admin/refunds/{case_id}/approve
│       │
│       └── reject/
│           └── page.tsx                    # Reject refund
│                                            # BE: admin-be/refund_case
│                                            # POST /v1/admin/refunds/{case_id}/reject
│
├── change-approvals/
│   ├── page.tsx                            # Change approval queue (Two-Person Rule)
│   │                                        # - Pending approvals
│   │                                        # - My requests
│   │                                        # - History
│   │                                        # BE: admin-be/change_approval
│   │                                        # GET /v1/admin/change-approvals
│   │
│   └── [requestId]/
│       ├── page.tsx                        # Change request detail
│       │                                    # BE: admin-be/change_approval
│       │                                    # GET /v1/admin/change-approvals/{request_id}
│       │
│       └── approve/
│           └── page.tsx                    # Approve/reject change
│                                            # BE: admin-be/change_approval
│                                            # POST /v1/admin/change-approvals/{request_id}/approve
│                                            # POST /v1/admin/change-approvals/{request_id}/reject
│
├── reports/
│   ├── page.tsx                            # Admin reports
│   │                                        # - Platform metrics
│   │                                        # - User growth
│   │                                        # - Revenue reports
│   │                                        # - Moderation stats
│   │                                        # BE: admin-be/reports
│   │                                        # GET /v1/admin/reports
│   │
│   ├── users/
│   │   └── page.tsx                        # User reports
│   │                                        # BE: admin-be/reports
│   │                                        # GET /v1/admin/reports/users
│   │
│   ├── financial/
│   │   └── page.tsx                        # Financial reports
│   │                                        # BE: admin-be/reports
│   │                                        # GET /v1/admin/reports/financial
│   │
│   └── moderation/
│       └── page.tsx                        # Moderation reports
│                                            # BE: admin-be/reports
│                                            # GET /v1/admin/reports/moderation
│
├── system/
│   ├── page.tsx                            # System settings
│   │                                        # - Feature flags
│   │                                        # - System configuration
│   │                                        # BE: admin-be/system
│   │                                        # GET /v1/admin/system/config
│   │
│   ├── feature-flags/
│   │   ├── page.tsx                        # Feature flags management
│   │   │                                    # - List flags
│   │   │                                    # - Toggle flags
│   │   │                                    # - Rollout percentage
│   │   │                                    # BE: subscriptions-be/feature_toggles
│   │   │                                    # GET /v1/admin/feature-flags
│   │   │
│   │   └── [flagId]/
│   │       └── edit/
│   │           └── page.tsx                # Edit feature flag
│   │                                        # BE: subscriptions-be/feature_toggles
│   │                                        # PUT /v1/admin/feature-flags/{flag_id}
│   │
│   ├── announcements/
│   │   ├── page.tsx                        # System announcements
│   │   │                                    # BE: communications-be/announcements
│   │   │                                    # GET /v1/admin/announcements
│   │   │
│   │   └── create/
│   │       └── page.tsx                    # Create announcement
│   │                                        # BE: communications-be/announcements
│   │                                        # POST /v1/admin/announcements
│   │
│   └── maintenance/
│       └── page.tsx                        # Maintenance mode
│                                            # - Enable/disable maintenance
│                                            # - Maintenance message
│                                            # BE: admin-be/system
│                                            # POST /v1/admin/system/maintenance
│
└── audit-logs/
    └── page.tsx                            # Audit logs
                                            # - All admin actions
                                            # - Filter by admin
                                            # - Filter by action type
                                            # - Search logs
                                            # BE: admin-be/audit
                                            # GET /v1/admin/audit-logs
```

---

**End of Part 4**

**Continue to Part 5 for:**
- Mobile App Structure
- Shared Packages (UI, Shared, Types, Config)
- Feature Modules Structure
- Components Library
