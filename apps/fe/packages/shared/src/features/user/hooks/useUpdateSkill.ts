import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { UpdateSkillRequest, FreelancerSkill } from '@skillsier/types';

interface UpdateSkillVariables {
  skillId: string;
  data: UpdateSkillRequest;
}

export function useUpdateSkill() {
  const queryClient = useQueryClient();

  return useMutation<FreelancerSkill, Error, UpdateSkillVariables>({
    mutationFn: ({ skillId, data }) => userApi.updateSkill(skillId, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.skills });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.freelancerProfile });
    },
  });
}