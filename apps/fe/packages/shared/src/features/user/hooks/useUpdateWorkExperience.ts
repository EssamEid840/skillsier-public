import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { UpdateWorkExperienceRequest, WorkExperience } from '@skillsier/types';

interface UpdateWorkExperienceVariables {
  experienceId: string;
  data: UpdateWorkExperienceRequest;
}

export function useUpdateWorkExperience() {
  const queryClient = useQueryClient();

  return useMutation<WorkExperience, Error, UpdateWorkExperienceVariables>({
    mutationFn: ({ experienceId, data }) => userApi.updateWorkExperience(experienceId, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.workExperience });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.freelancerProfile });
    },
  });
}