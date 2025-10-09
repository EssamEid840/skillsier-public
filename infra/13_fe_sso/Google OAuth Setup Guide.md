# Google OAuth Setup Guide for Skillsier Keycloak

This guide will help you set up Google OAuth integration for your Keycloak instance to enable Google Sign-in/Sign-up functionality.

## 🔧 Step 1: Google Cloud Console Setup

### 1.1 Create/Select a Project
1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project or select an existing one
3. Name it something like "Skillsier Authentication"

### 1.2 Enable Required APIs
1. Go to **APIs & Services** > **Library**
2. Search for and enable:
   - **Google+ API** (for user profile info)
   - **People API** (recommended for better user data)

### 1.3 Configure OAuth Consent Screen
1. Go to **APIs & Services** > **OAuth consent screen**
2. Choose **External** user type (unless you have Google Workspace)
3. Fill in the required information:
   - **App name**: Skillsier
   - **User support email**: Your email
   - **App domain**: skillsier.com
   - **Authorized domains**: Add `skillsier.com`
   - **Developer contact information**: Your email
4. Add scopes:
   - `openid`
   - `email`
   - `profile`
5. Save and continue

### 1.4 Create OAuth 2.0 Credentials
1. Go to **APIs & Services** > **Credentials**
2. Click **+ CREATE CREDENTIALS** > **OAuth 2.0 Client IDs**
3. Choose **Web application**
4. Name it "Skillsier Keycloak"
5. Add **Authorized redirect URIs**:
   ```
   https://keycloak.skillsier.com/realms/skillsier/broker/google/endpoint
   ```
6. Click **Create**
7. **Save the Client ID and Client Secret** - you'll need these!

## 🚀 Step 2: Keycloak Configuration

### 2.1 Automatic Configuration (Recommended)
Run the configuration script and choose to configure Google OAuth:
```bash
./configure-keycloak-prod.sh
```

When prompted, enter:
- Your Google OAuth Client ID
- Your Google OAuth Client Secret

### 2.2 Manual Configuration (Alternative)
If you prefer to configure manually:

1. Go to Keycloak Admin Console: `https://keycloak.skillsier.com/admin/`
2. Select the **skillsier** realm
3. Go to **Identity Providers**
4. Click **Add provider** > **Google**
5. Fill in:
   - **Client ID**: Your Google OAuth Client ID
   - **Client Secret**: Your Google OAuth Client Secret
   - **Display Name**: Google
   - **Trust Email**: ON
   - **First Broker Login Flow**: first broker login
6. Save

### 2.3 Configure Attribute Mappers
The script automatically creates these mappers, but if configuring manually:

1. In the Google identity provider settings, go to **Mappers**
2. Create these mappers:

**Email Mapper:**
- Name: `google-email-mapper`
- Mapper Type: `Attribute Importer`
- Claim: `email`
- User Attribute Name: `email`

**First Name Mapper:**
- Name: `google-firstname-mapper`
- Mapper Type: `Attribute Importer`
- Claim: `given_name`
- User Attribute Name: `firstName`

**Last Name Mapper:**
- Name: `google-lastname-mapper`
- Mapper Type: `Attribute Importer`
- Claim: `family_name`
- User Attribute Name: `lastName`

## 🧪 Step 3: Testing

### 3.1 Test the Login Flow
1. Go to your application's login page
2. You should see a "Sign in with Google" button
3. Click it and test the Google OAuth flow
4. Verify that user information is correctly imported

### 3.2 Direct Test URL
You can also test directly via:
```
https://keycloak.skillsier.com/realms/skillsier/protocol/openid-connect/auth?client_id=skillsier-web&redirect_uri=https://skillsier.com&response_type=code&scope=openid
```

## 🔒 Security Considerations

### Domain Restrictions
To restrict Google OAuth to specific domains:
1. In Google Cloud Console OAuth consent screen
2. Add only your authorized domains
3. In Keycloak Google provider config, set **Hosted Domain** if needed

### User Roles
By default, users signing up via Google will get the basic "user" role. To automatically assign roles:
1. Go to **Identity Providers** > **Google** > **Mappers**
2. Create a **Hardcoded Role** mapper
3. Select the role you want to assign (e.g., "student")

## 🎯 Integration with Your Applications

### Frontend Integration
Your applications can now redirect users to:
```javascript
const googleSignInUrl = `https://keycloak.skillsier.com/realms/skillsier/protocol/openid-connect/auth?client_id=skillsier-web&redirect_uri=${encodeURIComponent(window.location.origin)}&response_type=code&scope=openid&kc_idp_hint=google`;

// Redirect directly to Google OAuth
window.location.href = googleSignInUrl;
```

### Backend Verification
Users authenticated via Google will have standard Keycloak JWT tokens that your `users-be` service can verify normally.

## 📋 Troubleshooting

### Common Issues:

1. **"redirect_uri_mismatch" error**
   - Ensure the redirect URI in Google Console exactly matches Keycloak's broker endpoint
   - Check for trailing slashes, http vs https

2. **"OAuth consent screen verification required"**
   - Your app needs verification if requesting sensitive scopes
   - For basic email/profile, this usually isn't required

3. **Users can't see Google sign-in option**
   - Check that the identity provider is enabled
   - Verify the realm is correct
   - Check browser console for errors

4. **User information not imported**
   - Verify the attribute mappers are configured correctly
   - Check that the Google API scopes include `profile` and `email`

## 🎉 What Users Will See

After setup, users will see:
- A "Sign in with Google" button on the Keycloak login page
- Seamless redirect to Google for authentication
- Automatic account creation for new users
- Pre-filled profile information from Google

This integration provides a smooth user experience while maintaining security through Keycloak's robust authentication system.