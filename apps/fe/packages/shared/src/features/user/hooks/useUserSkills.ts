import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { AddSkillRequest, UpdateSkillRequest } from '@skillsier/types';

export const useUserSkills = () => {
  return useQuery({
    queryKey: queryKeys.users.skills,
    queryFn: userApi.getSkills,
  });
};

export const useAddSkill = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: AddSkillRequest) => userApi.addSkill(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.skills });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.profile });
    },
  });
};

export const useUpdateSkill = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ skillId, data }: { skillId: string; data: UpdateSkillRequest }) =>
      userApi.updateSkill(skillId, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.skills });
    },
  });
};

export const useDeleteSkill = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (skillId: string) => userApi.deleteSkill(skillId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.skills });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.profile });
    },
  });
};