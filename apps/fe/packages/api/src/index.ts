export { BaseApiClient, ApiError } from './lib/base-client';
export type { ApiClientConfig } from './lib/base-client';
export {
  handleApiError,
  isApiError,
  getErrorMessage,
} from './lib/error-handler';

export { JobsClient } from './clients/jobs-client';
export { UsersClient } from './clients/users-client';
export { ProposalsClient } from './clients/proposals-client';

export { MockJobsClient } from './mocks/jobs-client.mock';
export { mockJobs, mockJobListResponse, getMockJob } from './mocks/jobs.mock';
export { mockUsers, mockUserProfile, getMockUser } from './mocks/users.mock';