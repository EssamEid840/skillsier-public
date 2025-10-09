import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { UpdateEducationRequest, Education } from '@skillsier/types';

interface UpdateEducationVariables {
  educationId: string;
  data: UpdateEducationRequest;
}

export function useUpdateEducation() {
  const queryClient = useQueryClient();

  return useMutation<Education, Error, UpdateEducationVariables>({
    mutationFn: ({ educationId, data }) => userApi.updateEducation(educationId, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.education });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.freelancerProfile });
    },
  });
}