# what is the differences between models and domains

 the key differences between **models** and **domains** in your frontend folder structure:

**📁 DOMAINS vs MODELS - Key Differences**
------------------------------------------

### **1\. DOMAINS** (packages/shared/src/types/domains/)

**Purpose:** Business logic organization - groups related types by business capability

**Characteristics:**

*   **Business-focused:** Organized by business domains (financial, contracts, jobs, etc.)
    
*   **Comprehensive:** Contains ALL types related to that domain
    
*   **Hierarchical:** Can have subfolders for complex entities
    
*   **Domain-driven:** Follows DDD (Domain-Driven Design) principles
    

**Example Structure:**

typescript


```

packages/shared/src/types/domains/
├── financial/
│   ├── wallet.ts              # Wallet domain types
│   ├── transaction.ts         # Transaction domain types
│   ├── payment.ts             # Payment domain types
│   ├── escrow.ts              # Escrow domain types
│   ├── invoice.ts             # Invoice domain types
│   └── index.ts               # Barrel export
```
**Content Example (financial/wallet.ts):**

typescript


```
// Domain types - comprehensive business logic representation
export interface Wallet {
  id: string;
  userId: string;
  balance: number;
  currency: string;
  status: WalletStatus;
  availableBalance: number;
  pendingBalance: number;
  reservedBalance: number;
  limits: WalletLimits;
  metadata: Record<string, any>;
  createdAt: Date;
  updatedAt: Date;
}

export interface WalletLimits {
  dailyWithdrawal: number;
  monthlyWithdrawal: number;
  minimumBalance: number;
}

export interface WalletTransaction {
  id: string;
  walletId: string;
  amount: number;
  type: TransactionType;
  status: TransactionStatus;
  reference: string;
  description: string;
}

export enum WalletStatus {
  ACTIVE = 'ACTIVE',
  FROZEN = 'FROZEN',
  SUSPENDED = 'SUSPENDED',
  CLOSED = 'CLOSED'
}

```
### **2\. MODELS** (packages/shared/src/types/models/)

**Purpose:** Data Transfer Objects (DTOs) - API request/response shapes

**Characteristics:**

*   **API-focused:** Organized by backend microservice
    
*   **Transport layer:** Data as it moves between frontend and backend
    
*   **Request/Response:** Often split into separate types for input/output
    
*   **Service-oriented:** Mirrors backend API structure
    

**Example Structure:**

typescript


```
packages/shared/src/types/models/
├── admin/
│   ├── admin-user.ts          # Admin user DTO
│   ├── support-ticket.ts      # Support ticket DTO
│   └── moderation-item.ts     # Moderation item DTO
├── financial/
│   ├── wallet-dto.ts          # Wallet API DTO
│   ├── payment-dto.ts         # Payment API DTO
│   └── transaction-dto.ts     # Transaction API DTO
```
**Content Example (models/financial/wallet-dto.ts):**

typescript


```
// DTOs - API request/response shapes
export interface WalletDTO {
  id: string;
  user_id: string;           // snake_case from API
  balance: string;           // String for precision
  currency: string;
  status: string;
  available_balance: string;
  created_at: string;        // ISO string from API
  updated_at: string;
}

export interface CreateWalletRequest {
  user_id: string;
  currency: string;
  initial_balance?: string;
}

export interface CreateWalletResponse {
  wallet: WalletDTO;
  message: string;
}

export interface UpdateWalletRequest {
  status?: string;
  limits?: {
    daily_withdrawal?: string;
    monthly_withdrawal?: string;
  };
}

export interface GetWalletBalanceResponse {
  balance: string;
  available_balance: string;
  pending_balance: string;
  reserved_balance: string;
  currency: string;
}
```

---

## **🔄 How They Work Together**

### **Data Flow:**
```
Backend API → Models (DTOs) → Domain Mappers → Domains (Business Types) → UI Components

```
### **Example Transformation:**

**1\. API Response (Model/DTO):**

typescript


```
// From backend - financial-be API
const walletDTO: WalletDTO = {
  id: "wallet_123",
  user_id: "user_456",
  balance: "1000.50",
  currency: "USD",
  status: "ACTIVE",
  available_balance: "800.00",
  created_at: "2025-01-15T10:30:00Z",
  updated_at: "2025-01-20T15:45:00Z"
};

```
**2\. Transform to Domain Type:**

typescript


```
// Mapper function
function walletDTOToDomain(dto: WalletDTO): Wallet {
  return {
    id: dto.id,
    userId: dto.user_id,              // camelCase conversion
    balance: parseFloat(dto.balance),  // Number conversion
    currency: dto.currency,
    status: dto.status as WalletStatus,
    availableBalance: parseFloat(dto.available_balance),
    pendingBalance: 0,                 // Calculated
    reservedBalance: 200.50,           // Calculated
    limits: {                          // Enriched
      dailyWithdrawal: 5000,
      monthlyWithdrawal: 50000,
      minimumBalance: 0
    },
    metadata: {},
    createdAt: new Date(dto.created_at),  // Date object
    updatedAt: new Date(dto.updated_at)
  };
}

```

**3\. Use in Component:**

typescript


```
// Component uses domain type
function WalletCard({ wallet }: { wallet: Wallet }) {
  return (
    <div>
      <h3>Balance: ${wallet.availableBalance.toFixed(2)}</h3>
      <p>Status: {wallet.status}</p>
      <p>Daily Limit: ${wallet.limits.dailyWithdrawal}</p>
    </div>
  );
}
```

---

## **📊 Comparison Table**

| Aspect | **DOMAINS** | **MODELS (DTOs)** |
|--------|------------|-------------------|
| **Purpose** | Business logic representation | API data transfer |
| **Organization** | By business capability | By backend service |
| **Naming** | Business terms (camelCase) | API terms (snake_case) |
| **Types** | Rich TypeScript types | Simple API shapes |
| **Usage** | Internal app logic, UI | API calls only |
| **Transformations** | Computed properties, enriched data | Raw API data |
| **Location** | `types/domains/` | `types/models/` |
| **Example File** | `domains/financial/wallet.ts` | `models/financial/wallet-dto.ts` |

---

## **🎯 When to Use Each**

### **Use DOMAINS when:**
- ✅ Rendering UI components
- ✅ Business logic calculations
- ✅ State management (Redux, Zustand)
- ✅ Form validation
- ✅ Client-side data manipulation
- ✅ Type checking in components

### **Use MODELS (DTOs) when:**
- ✅ Making API calls
- ✅ Receiving API responses
- ✅ Sending API requests
- ✅ Defining API client interfaces
- ✅ Validating API payloads
- ✅ Type checking API layer

---

## **📂 Your Current Structure**

Based on your `wo-comments-combined-fe-folder-strucure.md`:
```
packages/shared/src/types/
├── api/                    # ❌ Backend API response types (should be in models/)
│   └── backend/
│       ├── admin-be.ts
│       ├── financial-be.ts
│       └── ...
├── domains/                # ✅ Business domain types
│   ├── admin/
│   ├── financial/
│   ├── contracts/
│   └── ...
├── entities/               # ⚠️ Similar to domains (might be redundant)
│   ├── admin/
│   ├── financial/
│   └── ...
├── enums/                  # ✅ Shared enumerations
│   ├── admin/
│   ├── financial/
│   └── ...
└── models/                 # ✅ DTOs (data transfer objects)
    ├── admin/
    ├── financial/
    └── ...
```
**💡 Recommendation: Clean Architecture**
-----------------------------------------

### **Suggested Organization:**

typescript


```
packages/shared/src/types/
├── api/                           # DELETE - move to models/
├── domains/                       # ✅ Keep - Business types
│   ├── financial/
│   │   ├── wallet.ts             # Wallet business type
│   │   ├── transaction.ts
│   │   └── index.ts
│   └── contracts/
│       ├── contract.ts
│       └── index.ts
├── entities/                      # ⚠️ MERGE into domains or keep for shared entities
│   ├── contract.ts               # Could be in domains/contracts/
│   ├── invoice.ts
│   └── user.ts
├── enums/                         # ✅ Keep - Shared enums
│   ├── financial/
│   │   ├── wallet-status.enum.ts
│   │   └── transaction-type.enum.ts
│   └── contracts/
├── models/                        # ✅ Keep - API DTOs
│   ├── admin/                    # Admin-BE API types
│   │   ├── admin-user.ts        # AdminUserDTO, CreateAdminUserRequest, etc.
│   │   └── support-ticket.ts
│   ├── financial/                # Financial-BE API types
│   │   ├── wallet-dto.ts        # WalletDTO, CreateWalletRequest, etc.
│   │   └── payment-dto.ts
│   └── contracts/                # Contracts-BE API types
└── mappers/                       # ✅ NEW - Transform DTOs to domains
    ├── financial/
    │   └── wallet.mapper.ts      # walletDTOToDomain(), walletToDTO()
    └── contracts/

```

**🔧 Practical Example: Complete Flow**
---------------------------------------

### **1\. Define DTO (Model):**

typescript


```
// packages/shared/src/types/models/financial/wallet-dto.ts
export interface WalletDTO {
  id: string;
  user_id: string;
  balance: string;
  currency: string;
  status: string;
}
```
### **2\. Define Domain Type:**

typescript


```
// packages/shared/src/types/domains/financial/wallet.ts
export interface Wallet {
  id: string;
  userId: string;
  balance: number;
  currency: string;
  status: WalletStatus;
  formattedBalance: string;  // Computed
}

export enum WalletStatus {
  ACTIVE = 'ACTIVE',
  FROZEN = 'FROZEN'
}

```
### **3\. Create Mapper:**

typescript


```
// packages/shared/src/types/mappers/financial/wallet.mapper.ts
export function walletDTOToDomain(dto: WalletDTO): Wallet {
  const balance = parseFloat(dto.balance);
  return {
    id: dto.id,
    userId: dto.user_id,
    balance,
    currency: dto.currency,
    status: dto.status as WalletStatus,
    formattedBalance: `${dto.currency} ${balance.toFixed(2)}`
  };
}

```

### **4\. Use in API Client:**

typescript


```

// apps/web/src/lib/api/financial/wallets.ts
export async function getWallet(id: string): Promise<Wallet> {
  const dto = await api.get<WalletDTO>(`/wallets/${id}`);
  return walletDTOToDomain(dto);  // Transform to domain
}

```
### **5\. Use in Component:**

typescript


```
// apps/web/src/components/WalletCard.tsx
function WalletCard() {
  const { data: wallet } = useQuery({
    queryKey: ['wallet'],
    queryFn: () => getWallet('wallet_123')
  });

  return <div>{wallet?.formattedBalance}</div>;
}

```
**Summary**
-----------

*   **DOMAINS** = Business logic types (internal use, rich types)
    
*   **MODELS** = API DTOs (external interface, raw data)
    
*   **ENTITIES** = Could be either (evaluate if needed)
    
*   **MAPPERS** = Transform between DTOs and domains
    

The key is **separation of concerns**: API layer uses DTOs, application layer uses domains, and mappers bridge the gap.