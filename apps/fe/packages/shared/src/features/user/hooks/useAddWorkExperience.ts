import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { AddWorkExperienceRequest, WorkExperience } from '@skillsier/types';

export function useAddWorkExperience() {
  const queryClient = useQueryClient();

  return useMutation<WorkExperience, Error, AddWorkExperienceRequest>({
    mutationFn: userApi.addWorkExperience,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.workExperience });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.freelancerProfile });
    },
  });
}