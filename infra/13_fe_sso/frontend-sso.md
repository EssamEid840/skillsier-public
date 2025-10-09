Awesome — since this is a **brand-new Keycloak**, here’s a clean, exact checklist to get **Login/Register with Google** working with your existing frontend.

I’ll assume:

*   Realm: skillsier
    
*   Keycloak base: https://keycloak.skillsier.com
    
*   Frontend: https://skillsier.com (and http://localhost:3000 for dev)
    
*   Your frontend already calls Keycloak with kc\_idp\_hint=google and expects the **frontend callback**/api/auth/keycloak/oauth/callback/google
    

If any of those differ, replace the names/URLs accordingly.

A) Create/verify the Realm
==========================

1.  Log in to **Keycloak Admin** → top-left realm selector → **Create realm**
    
    *   **Name:** skillsier
        
2.  **Realm Settings → General**
    
    *   **Frontend URL:** https://keycloak.skillsier.com/
        
    *   **Require SSL:** all external requests (recommended)
        
3.  **Realm Settings → Login** (optional but recommended)
    
    *   **Forgot password:** ON (helps local accounts)
        
    *   **Verify email:** ON(turn **Trust Email** ON in the Google IdP below so Google users skip extra verification)
        
    *   **User registration:** ON only if you want form-based sign-ups in addition to Google
        

> These do not block Google SSO; Google SSO works even if “User registration” is OFF.

B) Create the **frontend OIDC client** (your web app)
=====================================================

1.  **Clients → Create client**
    
    *   **Client type:** OpenID Connect
        
    *   **Client ID:** skillsier-fe
        
    *   Next → **Capability config**:
        
        *   **Standard flow**: ON
            
        *   **Direct access grants**: OFF
            
        *   **Service accounts**: OFF
            
        *   **Client authentication**: **ON** (confidential client, since your Next.js API exchanges the code)
            
        *   **PKCE**: **S256** (recommended)
            
    *   Create
        
2.  **Clients → skillsier-fe → Settings**
    
    *   http://localhost:3000/api/auth/keycloak/oauth/callback/googlehttps://skillsier.com/api/auth/keycloak/oauth/callback/google
        
    *   http://localhost:3000https://skillsier.com
        
    *   http://localhost:3000/\*https://skillsier.com/\*
        
    *   Save
        
3.  **Clients → skillsier-fe → Credentials**
    
    *   Copy the **Client secret** and put it in your FE env as KEYCLOAK\_CLIENT\_SECRET.
        

> The built-in **default client scopes** (profile, email, roles) should already be attached. If not, attach profile and email in **Clients → skillsier-fe → Client scopes** (Default Client Scopes tab).

C) Create the **management (BFF) client** for registration APIs
===============================================================

Your backend/Next.js API uses client-credentials to call Keycloak Admin (to create/link users during register).

1.  **Clients → Create client**
    
    *   **Client type:** OpenID Connect
        
    *   **Client ID:** skillsier-bff
        
    *   Next → **Capability config**:
        
        *   **Standard flow**: OFF
            
        *   **Client authentication**: ON
            
        *   **Service accounts**: **ON**
            
    *   Create
        
2.  **Clients → skillsier-bff → Service account roles**
    
    *   **Client roles → realm-management**:
        
        *   Assign: manage-users, view-users
            
        *   (Optionally: view-realm, manage-account-links if you plan to auto-link or do email actions)
            
3.  **Clients → skillsier-bff → Credentials**
    
    *   Copy the **Client secret** → set in your FE env as KEYCLOAK\_MGMT\_CLIENT\_SECRET.
        

D) Add **Google** as an Identity Provider (broker)
==================================================

> You must already have the **Google OAuth Client ID/Secret** created in Google Cloud with this **Authorized redirect URI**:
> 

 https://keycloak.skillsier.com/realms/skillsier/broker/google/endpoint   `

1.  **Identity Providers → Add provider → Google**
    
    *   **Alias:** google ← keep exactly google (your FE hints with kc\_idp\_hint=google)
        
    *   **Client ID / Client Secret:** from Google
        
    *   **Default Scopes / Scopes:** openid email profile
        
    *   **Trust Email:** **ON**
        
    *   **Store Tokens:** OFF (or ON if you need downstream Google tokens)
        
    *   **Sync Mode:** Import
        
    *   **First Broker Login Flow:** first broker login (default)
        
    *   **Hide on login page:** OFF (so it shows a “Google” button on Keycloak’s login page too)
        
    *   Save
        
2.  https://keycloak.skillsier.com/realms/skillsier/broker/google/endpoint
    
3.  **(Optional) Identity Provider → google → Mappers → Add mapper**
    
    *   Type: **Attribute Importer**
        
        *   Claim: email → **User Attribute**: email (Update existing: ON)
            
        *   Claim: given\_name → **User Attribute**: firstName
            
        *   Claim: family\_name → **User Attribute**: lastName
            
        *   Claim: picture → **User Attribute**: picture (if you want avatars)
            

> Default mappers usually cover names; adding them explicitly avoids surprises.

E) Optional: streamline first-login behavior
============================================

If you want **silent auto-creation** (no “review profile” screen):

1.  **Authentication → Flows → First Broker Login**
    
    *   Click **Actions → Duplicate** (e.g., first broker login (no review))
        
    *   In the duplicated flow, **Disable** the “Review Profile” execution (or set it to Alternative/Not required).
        
    *   Ensure “Create User If Unique” is **Required**.
        
2.  Go back to **Identity Providers → google → First Broker Login Flow** → select your new flow.
    

F) Sanity settings that affect redirects
========================================

*   **Realm Settings → General → Frontend URL:** https://keycloak.skillsier.com/
    
*   Your ingress should set X-Forwarded-Proto: https so Keycloak generates **https** links.
    
*   If you still see http:// redirects, set **Hostname**/proxy arguments in your container (which you already do) and keep **Frontend URL** set.
    

G) Quick end-to-end test
========================

1.  In **Identity Providers → google**, click **Initiate Login**.You should be sent to Google, then back to Keycloak, then back to your FE callback.
    
2.  From your app, hit your Google button (it sends kc\_idp\_hint=google to Keycloak).On success, you should land on /users.
    

Copy-paste recap (values to use)
--------------------------------

**Keycloak → Clients → skillsier-fe**

* Redirects:
- htp://localhost:3000/api/auth/keycloak/oauth/callback/google
- htps://skillsier.com/api/auth/keycloak/oauth/callback/google
    
* Web origins:  
- http://localhost:3000
- https://skillsier.com
    
*  Post-logout redirects: 
- http://localhost:3000/*
- https://skillsier.com/*
    

**Keycloak → Identity Providers → Google**

```text  

Alias: google
Client ID / Secret: (from Google)
Scopes: openid email profile
Trust Email: ON
Redirect URI (verify): https://keycloak.skillsier.com/realms/skillsier/broker/google/endpoint


```


**Realm Settings**

```text  

Frontend URL: https://keycloak.skillsier.com/
Require SSL: all external requests


```
**Service client (skillsier-bff)**

```text  

Service accounts: ON
realm-management roles: manage-users, view-users (plus optional others)


```  
If your **realm name** isn’t skillsier or your **frontend domain** isn’t skillsier.com, tell me the exact values and I’ll re-emit this with your precise URIs and the two client configs ready to paste.



You’re getting:

```text  

error=invalid_request
error_description=Missing parameter: code_challenge_method

```  

That means your **Keycloak client requires PKCE**, but your auth request (from /api/auth/keycloak/oauth/start/google) isn’t sending code\_challenge + code\_challenge\_method. You have two ways to fix this. I’ll give you the **fast fix (no PKCE)** and the **recommended fix (keep PKCE)**, plus one more Keycloak check for your wrong iss value.

✅ Fast fix (do this now to get unblocked)
-----------------------------------------

In **Keycloak Admin** for your frontend client (e.g., skillsier-fe):

1.  **Clients → skillsier-fe → Settings**
    
    *   Ensure **Standard flow** = ON
        
    *   Ensure **Client authentication** = ON (confidential client)
        
2.  **Clients → skillsier-fe → Advanced** (tab may be “Capabilities/Advanced” depending on theme)
    
    *   **PKCE**/**Proof Key for Code Exchange**: set to **Not required** (or set “Code challenge method” to **None**).
        
    *   **Save**.
        

3) In Keycloak → Realm Settings → General
-----------------------------------------

*   **Frontend URL:** https://keycloak.skillsier.com/
    

4) In Keycloak → Identity Providers → Google
--------------------------------------------

*   **Alias:** google
    
*   **Client ID/Secret:** from Google Console
    
*   **Scopes:** openid email profile
    
*   **Trust Email:** ON



























































No code changes needed. Try login again.

🔒 Recommended (keep PKCE = S256) — if you want stronger security
-----------------------------------------------------------------

Keep **PKCE required (S256)** in Keycloak, but then your **start** route must add PKCE params and your **callback** route must send the code\_verifier when exchanging the code.

### Keycloak settings (if keeping PKCE)

*   **Clients → skillsier-fe → Advanced**
    
    *   **PKCE policy**: **Required**
        
    *   **Code challenge method**: **S256**
        
    *   Save.
        

### App behavior needed (summary)

*   ...&code\_challenge=&code\_challenge\_method=S256Store code\_verifier (encrypted cookie/session).
    
*   grant\_type=authorization\_codecode=redirect\_uri=client\_id=skillsier-feclient\_secret= # because your client is confidentialcode\_verifier=
    

> If you want, I’ll drop a minimal diff for your two routes that matches your current structure.

One more Keycloak-side fix for your logs
----------------------------------------

Your error line shows:

iss=http://localhost:3000/realms/skillsier 

This indicates your app is pointing KEYCLOAK\_ISSUER\_URL at **localhost**. Make sure your app uses the real Keycloak URL:

*   KEYCLOAK\_ISSUER\_URL=https://keycloak.skillsier.com/realms/skillsier(If you actually run Keycloak locally for dev, point to that Keycloak host:port. But **never** localhost:3000 — that’s your Next.js, not Keycloak.)
    

Also in **Keycloak**:

*   **Realm Settings → General**
    
    *   **Frontend URL:** https://keycloak.skillsier.com/
        
    *   **Require SSL:** All external requests
        

This keeps Keycloak generating **https** redirects.

Final Keycloak checklist (everything you need set)
--------------------------------------------------

**Realm (skillsier)**

*   Realm Settings → General → **Frontend URL** = https://keycloak.skillsier.com/
    

**Identity Provider (Google)**

*   Identity Providers → **Google**
    
    *   Alias: google
        
    *   Client ID/Secret: from Google console
        
    *   Scopes: openid email profile
        
    *   Trust Email: **ON**
        
    *   https://keycloak.skillsier.com/realms/skillsier/broker/google/endpoint
        

**Frontend Client (skillsier-fe)**

*   Clients → skillsier-fe → Settings
    
    *   Standard flow: **ON**
        
    *   Client authentication: **ON**
        
    *   http://localhost:3000/api/auth/keycloak/oauth/callback/googlehttps://skillsier.com/api/auth/keycloak/oauth/callback/google
        
    *   http://localhost:3000https://skillsier.com
        
    *   http://localhost:3000/\*https://skillsier.com/\*
        
*   Clients → skillsier-fe → Advanced
    
    *   **Fast fix**: PKCE **Not required** (or Code challenge method **None**)
        
    *   **Recommended**: PKCE **Required** + **S256** (but then implement PKCE in your start/callback)
        

**Management/BFF Client (skillsier-bff)**

*   Clients → skillsier-bff
    
    *   Service accounts: **ON**
        
    *   Service account roles (realm-management): manage-users, view-users (and any you need)
        

Try the **fast fix** first. If it works (it should), tell me if you want the small PKCE code snippet for your two routes and I’ll tailor it to your exact files.