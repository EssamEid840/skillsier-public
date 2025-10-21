
4) Email Delivery — failure path
================================

### 17.1.x delivery/ (Send Failure Handling → email.failed.v1)

##### Stories

*   As a **system**, I want to **emit an email failed event** when rendering or SMTP delivery definitively fails so diagnostics are clear.
    
*   As an **operator**, I want to **distinguish permanent vs transient** failures to drive retries or suppression.
    

##### Flow

1.  **SendEmailCommand**(user\_id, template\_id, template\_data, category, idempotency\_key)→ ValidateUser() | CheckPreferences() | CheckUnsubscribe()| Try RenderTemplate() → If RenderError(non-retryable)→ **Outbox:** email.failed.v1 {reason="render\_error", template\_id, error\_code, correlation\_id}| Else SendViaWildDuck()
    
    *   If SMTP 2xx → **Outbox:** email.sent.v1
        
    *   If SMTP 4xx/5xx transient → Queue **SendEmailRetryJob(email\_id, attempt+1)**
        
    *   If SMTP permanent (e.g., 550/5.1.1) → **Outbox:** email.failed.v1 {reason="smtp\_permanent", smtp\_code, smtp\_msg}
        
2.  **SendEmailRetryJob**(email\_id, attempt)→ ExponentialBackoff(2^n \* 100ms, max 5) | SendViaWildDuck()| If success → **Outbox:** email.sent.v1| If max\_retries\_exhausted → **Outbox:** email.failed.v1 {reason="retries\_exhausted", attempts=attempt}
    
3.  **(Optional) Bounce/Complaint handlers remain as-is**→ They continue to emit email.bounced.v1 / email.complaint.received.v1.
    

##### Projections

*   email\_delivery\_read (status=failed with reason)
    
*   email\_bounces\_read (unchanged)
    

##### Events Published

*   email.failed.v1 (render\_error | smtp\_permanent | retries\_exhausted)
    

##### RBAC/SLO

*   **RBAC:** SYSTEM (send/fail), OWNER (view status), ADMIN (view all stats)
    
*   **SLO:** Failure path P95 < 150ms (no retry); retry backoff capped; hand-off metrics preserved