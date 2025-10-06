/**
 * Public SDK surface for users + auth.
 * Import paths remain stable for web & mobile apps.
 */

/* auth */
export * as Auth from "../auth/login";
export * as Register from "../auth/register";
export * as Refresh from "../auth/refresh";
export * as Logout from "../auth/logout";
export * as Me from "../auth/me";
export * as VerifyEmail from "../auth/verify-email";
export * as ForgotPassword from "../auth/forgot-password";
export * as ResetPassword from "../auth/reset-password";
export * as GoogleProvider from "../auth/providers/startGoogle";
export * as GoogleCallback from "../auth/providers/callbackGoogle";

/* users REST */
export * as Profile from "./rest/profile";
export * as Portfolio from "./rest/portfolio";
export * as Verification from "./rest/verification";
export * as Contact from "./rest/contact";
export * as Connected from "./rest/connected";
export * as Sessions from "./rest/sessions";
export * as Security from "./rest/security";
export * as Notifications from "./rest/notifications";
export * as Preferences from "./rest/preferences";
export * as Privacy from "./rest/privacy";
export * as Billing from "./rest/billing";
export * as Developer from "./rest/developer";
export * as Referrals from "./rest/referrals";
export * as Onboarding from "./rest/onboarding";
export * as Uploads from "./rest/uploads";
export * as Abuse from "./rest/abuse";

/* users GraphQL */
export * as Gql from "./graphql";
