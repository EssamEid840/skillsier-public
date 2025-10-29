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
















**📁 STATE Types - Purpose & Usage**
------------------------------------

### **Location:**


```

packages/shared/src/types/state/
└── index.ts  # State management types

```
### **Purpose:**

State types define the **shape of your application state** in state management solutions (Redux, Zustand, Jotai, etc.)

**🎯 What Goes in state/**
--------------------------

### **1\. Store State Interfaces**

typescript


```
// packages/shared/src/types/state/index.ts

/**
 * Root application state
 */
export interface RootState {
  auth: AuthState;
  user: UserState;
  financial: FinancialState;
  contracts: ContractsState;
  jobs: JobsState;
  proposals: ProposalsState;
  ui: UIState;
}

/**
 * Auth slice state
 */
export interface AuthState {
  isAuthenticated: boolean;
  user: User | null;
  accessToken: string | null;
  refreshToken: string | null;
  loading: boolean;
  error: string | null;
}

/**
 * Financial slice state
 */
export interface FinancialState {
  wallets: Record<string, Wallet>;
  selectedWalletId: string | null;
  transactions: Transaction[];
  loading: boolean;
  error: string | null;
}

/**
 * UI slice state (modals, sidebars, etc.)
 */
export interface UIState {
  modals: {
    createWallet: boolean;
    sendPayment: boolean;
  };
  sidebar: {
    isOpen: boolean;
    activeSection: string | null;
  };
  theme: 'light' | 'dark';
}

```
### **2\. Action Types (Redux/Flux Pattern)**

typescript


```
// packages/shared/src/types/state/index.ts

/**
 * Base action interface
 */
export interface Action<T = any> {
  type: string;
  payload?: T;
  error?: boolean;
  meta?: Record<string, any>;
}

/**
 * Auth actions
 */
export type AuthAction =
  | { type: 'auth/login/pending' }
  | { type: 'auth/login/fulfilled'; payload: { user: User; tokens: Tokens } }
  | { type: 'auth/login/rejected'; payload: { error: string } }
  | { type: 'auth/logout' };

/**
 * Financial actions
 */
export type FinancialAction =
  | { type: 'financial/wallet/fetch/pending' }
  | { type: 'financial/wallet/fetch/fulfilled'; payload: Wallet }
  | { type: 'financial/wallet/fetch/rejected'; payload: { error: string } }
  | { type: 'financial/wallet/select'; payload: { walletId: string } }
  | { type: 'financial/transaction/add'; payload: Transaction };

```
### **3\. Reducer Types**

typescript


```
// packages/shared/src/types/state/index.ts

/**
 * Generic reducer type
 */
export type Reducer<S, A extends Action> = (state: S, action: A) => S;

/**
 * Auth reducer type
 */
export type AuthReducer = Reducer<AuthState, AuthAction>;

/**
 * Financial reducer type
 */
export type FinancialReducer = Reducer<FinancialState, FinancialAction>;

```
### **4\. Selector Types**

typescript


```
// packages/shared/src/types/state/index.ts

/**
 * Selector function type
 */
export type Selector<S, R> = (state: S) => R;

/**
 * Auth selectors
 */
export type AuthSelector<R> = Selector<RootState, R>;

/**
 * Example: Select current user
 */
export type SelectCurrentUser = AuthSelector<User | null>;

/**
 * Example: Select wallet by ID
 */
export type SelectWalletById = (state: RootState, walletId: string) => Wallet | undefined;
```
### **5\. Async Action Types (Thunks)**

typescript


```
// packages/shared/src/types/state/index.ts

/**
 * Thunk action type (for Redux Thunk)
 */
export type ThunkAction<R = void> = (
  dispatch: Dispatch,
  getState: () => RootState
) => R | Promise<R>;

/**
 * Async thunk type
 */
export type AsyncThunk<Returned = void, ThunkArg = void> = (
  arg: ThunkArg
) => ThunkAction<Promise<Returned>>;

/**
 * Example: Login thunk
 */
export type LoginThunk = AsyncThunk
  { user: User; tokens: Tokens },
  { email: string; password: string }
>;

```
### **6\. Store Configuration Types**

typescript


```
// packages/shared/src/types/state/index.ts

/**
 * Store configuration
 */
export interface StoreConfig {
  preloadedState?: Partial<RootState>;
  middleware?: Middleware[];
  devTools?: boolean;
}

/**
 * Middleware type
 */
export type Middleware = (
  api: MiddlewareAPI
) => (next: Dispatch) => (action: Action) => any;

/**
 * Middleware API
 */
export interface MiddlewareAPI {
  dispatch: Dispatch;
  getState: () => RootState;
}

/**
 * Dispatch type
 */
export type Dispatch = (action: Action | ThunkAction) => any;
```

---

## **🔄 How State Types Relate to Other Types**

### **Visual Hierarchy:**
```
┌─────────────────────────────────────────┐
│         APPLICATION LAYERS              │
├─────────────────────────────────────────┤
│                                         │
│  ┌──────────────────────────────────┐   │ 
│  │    UI COMPONENTS (React)         │   │
│  │  - Use: Domains & State types    │   │
│  └──────────────────────────────────┘   │
│               ↕                         │
│  ┌──────────────────────────────────┐   │
│  │    STATE MANAGEMENT              │   │
│  │  - Use: State types              │   │
│  │  - Store: Domain types           │   │
│  └──────────────────────────────────┘   │
│               ↕                         │
│  ┌──────────────────────────────────┐   │
│  │    MAPPERS                       │   │
│  │  - Transform: DTOs → Domains     │   │
│  └──────────────────────────────────┘   │
│               ↕                         │
│  ┌──────────────────────────────────┐   │
│  │    API CLIENT                    │   │
│  │  - Use: Models (DTOs)            │   │
│  └──────────────────────────────────┘   │
│               ↕                         │
│  ┌──────────────────────────────────┐   │
│  │    BACKEND API                   │   │
│  └──────────────────────────────────┘   │
│                                         │
└─────────────────────────────────────────┘

```
**📊 Comparison: State vs Domains vs Models**
---------------------------------------------


```
| Aspect   | STATE                    | DOMAINS            | MODELS (DTOs)        |
|----------|--------------------------|--------------------|----------------------|
| Purpose  | State management shape   | Business logic     | API transport        |
| Scope    | Application state tree   | Business entities  | API payloads         |
| Usage    | Redux/Zustand stores     | UI components      | API calls            |
| Contains | Store slices, actions    | Business types     | Request/response     |
| Example  | `RootState`, `AuthState` | `Wallet`, `Contract` | `WalletDTO`, `ContractDTO` |
| Location | `types/state/`           | `types/domains/`   | `types/models/`      |


```
**🛠️ Complete Example: Redux Implementation**
----------------------------------------------

### **1\. State Types**

typescript


```
// packages/shared/src/types/state/financial.state.ts

import { Wallet, Transaction } from '../domains/financial';

export interface FinancialState {
  wallets: {
    byId: Record<string, Wallet>;
    allIds: string[];
  };
  selectedWalletId: string | null;
  transactions: {
    byId: Record<string, Transaction>;
    allIds: string[];
  };
  loading: {
    wallets: boolean;
    transactions: boolean;
  };
  error: {
    wallets: string | null;
    transactions: string | null;
  };
}

export type FinancialAction =
  | { type: 'financial/wallets/fetch/pending' }
  | { type: 'financial/wallets/fetch/fulfilled'; payload: Wallet[] }
  | { type: 'financial/wallets/fetch/rejected'; payload: { error: string } }
  | { type: 'financial/wallet/select'; payload: { walletId: string } };

```
### **2\. Initial State**

```
// apps/web/src/store/slices/financial.slice.ts

import { FinancialState } from '@shared/types/state';

const initialState: FinancialState = {
  wallets: {
    byId: {},
    allIds: []
  },
  selectedWalletId: null,
  transactions: {
    byId: {},
    allIds: []
  },
  loading: {
    wallets: false,
    transactions: false
  },
  error: {
    wallets: null,
    transactions: null
  }
};
```
### **3\. Reducer**

typescript


```
// apps/web/src/store/slices/financial.slice.ts

import { FinancialState, FinancialAction } from '@shared/types/state';

export const financialReducer = (
  state: FinancialState = initialState,
  action: FinancialAction
): FinancialState => {
  switch (action.type) {
    case 'financial/wallets/fetch/pending':
      return {
        ...state,
        loading: { ...state.loading, wallets: true },
        error: { ...state.error, wallets: null }
      };
    
    case 'financial/wallets/fetch/fulfilled':
      return {
        ...state,
        wallets: {
          byId: action.payload.reduce((acc, wallet) => {
            acc[wallet.id] = wallet;
            return acc;
          }, {} as Record<string, Wallet>),
          allIds: action.payload.map(w => w.id)
        },
        loading: { ...state.loading, wallets: false }
      };
    
    case 'financial/wallets/fetch/rejected':
      return {
        ...state,
        loading: { ...state.loading, wallets: false },
        error: { ...state.error, wallets: action.payload.error }
      };
    
    case 'financial/wallet/select':
      return {
        ...state,
        selectedWalletId: action.payload.walletId
      };
    
    default:
      return state;
  }
};

```
### **4\. Selectors**

typescript


```
// apps/web/src/store/selectors/financial.selectors.ts

import { RootState } from '@shared/types/state';
import { Wallet } from '@shared/types/domains/financial';

export const selectAllWallets = (state: RootState): Wallet[] => {
  const { byId, allIds } = state.financial.wallets;
  return allIds.map(id => byId[id]);
};

export const selectWalletById = (
  state: RootState,
  walletId: string
): Wallet | undefined => {
  return state.financial.wallets.byId[walletId];
};

export const selectSelectedWallet = (state: RootState): Wallet | null => {
  const { selectedWalletId, wallets } = state.financial;
  return selectedWalletId ? wallets.byId[selectedWalletId] : null;
};

export const selectWalletsLoading = (state: RootState): boolean => {
  return state.financial.loading.wallets;
};

```
### **5\. Async Actions (Thunks)**

typescript


```
// apps/web/src/store/actions/financial.actions.ts

import { ThunkAction } from '@shared/types/state';
import { getWallets } from '@/lib/api/financial';

export const fetchWallets = (): ThunkAction => async (dispatch, getState) => {
  dispatch({ type: 'financial/wallets/fetch/pending' });
  
  try {
    // API call returns Domain type (already mapped from DTO)
    const wallets = await getWallets();
    
    dispatch({
      type: 'financial/wallets/fetch/fulfilled',
      payload: wallets
    });
  } catch (error) {
    dispatch({
      type: 'financial/wallets/fetch/rejected',
      payload: { error: error.message }
    });
  }
};

```
### **6\. Use in Component**

typescript


```
// apps/web/src/components/WalletList.tsx

import { useSelector, useDispatch } from 'react-redux';
import { selectAllWallets, selectWalletsLoading } from '@/store/selectors';
import { fetchWallets } from '@/store/actions';

export function WalletList() {
  const dispatch = useDispatch();
  const wallets = useSelector(selectAllWallets);     // Domain types
  const loading = useSelector(selectWalletsLoading); // State types
  
  useEffect(() => {
    dispatch(fetchWallets());
  }, []);
  
  if (loading) return <Spinner />;
  
  return (
    <div>
      {wallets.map(wallet => (
        <WalletCard key={wallet.id} wallet={wallet} />
      ))}
    </div>
  );
}
```
**🎯 Alternative: Zustand Example**
-----------------------------------

typescript


```
// packages/shared/src/types/state/financial.state.ts

import { Wallet } from '../domains/financial';

export interface FinancialStore {
  // State
  wallets: Wallet[];
  selectedWalletId: string | null;
  loading: boolean;
  error: string | null;
  
  // Actions
  fetchWallets: () => Promise<void>;
  selectWallet: (walletId: string) => void;
  clearError: () => void;
}

```
typescript


```
// apps/web/src/store/financial.store.ts

import { create } from 'zustand';
import { FinancialStore } from '@shared/types/state';
import { getWallets } from '@/lib/api/financial';

export const useFinancialStore = create<FinancialStore>((set) => ({
  wallets: [],
  selectedWalletId: null,
  loading: false,
  error: null,
  
  fetchWallets: async () => {
    set({ loading: true, error: null });
    try {
      const wallets = await getWallets();
      set({ wallets, loading: false });
    } catch (error) {
      set({ error: error.message, loading: false });
    }
  },
  
  selectWallet: (walletId) => {
    set({ selectedWalletId: walletId });
  },
  
  clearError: () => {
    set({ error: null });
  }
}));

```
**📁 Recommended Structure for State Types**
--------------------------------------------

typescript


```
packages/shared/src/types/state/
├── index.ts                          # Main exports
├── root.state.ts                     # RootState interface
├── actions.ts                        # Action types
├── reducers.ts                       # Reducer types
├── selectors.ts                      # Selector types
├── thunks.ts                         # Async action types
├── middleware.ts                     # Middleware types
├── slices/                           # State slice types
│   ├── auth.state.ts                 # AuthState, AuthAction
│   ├── financial.state.ts            # FinancialState, FinancialAction
│   ├── contracts.state.ts            # ContractsState, ContractsAction
│   ├── jobs.state.ts                 # JobsState, JobsAction
│   ├── proposals.state.ts            # ProposalsState, ProposalsAction
│   ├── ui.state.ts                   # UIState, UIAction
│   └── index.ts                      # Barrel export
└── stores/                           # Store-specific types (Zustand/Jotai)
    ├── financial.store.ts            # FinancialStore interface
    ├── contracts.store.ts            # ContractsStore interface
    └── index.ts                      # Barrel export
```

---

## **✨ Summary**

| Type | Purpose | Contains | Used By |
|------|---------|----------|---------|
| **State** | Application state shape | Store slices, actions, reducers | State management (Redux, Zustand) |
| **Domains** | Business entities | Business logic types | Components, state stores |
| **Models** | API transport | DTOs, requests, responses | API clients |
| **Entities** | Shared data structures | Basic entities | Domains, models |
| **Enums** | Constant values | Enumerations | All layers |

**Flow:**
```
API (Models/DTOs) → Mappers → Domains → State Store → Components

```
The **state types** define **how your application organizes and manages data in memory**, while **domains** define **what that data represents** and **models** define **how it's transferred**.
















































































































**Summary**
-----------

*   **DOMAINS** = Business logic types (internal use, rich types)
    
*   **MODELS** = API DTOs (external interface, raw data)
    
*   **ENTITIES** = Could be either (evaluate if needed)
    
*   **MAPPERS** = Transform between DTOs and domains
    

The key is **separation of concerns**: API layer uses DTOs, application layer uses domains, and mappers bridge the gap.