## Step 2: Create the DNS Records in Cloudflare
Now, log in to your Cloudflare dashboard and follow these steps.

Select your domain, skillsier.com.

Navigate to the DNS > Records section on the left sidebar.

Click Add record.

Record 1: Dapr Dashboard
Fill in the fields like this:

Type: A

Name: dapr (Cloudflare will automatically add skillsier.com)

IPv4 address: Paste the external IP address you found in Step 1.

Proxy status: Proxied (make sure the cloud icon is orange). This gives you free SSL and hides your server's IP.

TTL: Auto

➡️ Click Save.

-------

Record Details:

Type: A
Name: dapr-api (Cloudflare will automatically add .skillsier.com)
IPv4 address: 173.212.218.251:443 (your server's external IP)
Proxy status: Proxied (orange cloud icon)
TTL: Auto
