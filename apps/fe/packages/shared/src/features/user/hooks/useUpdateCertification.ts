import { useMutation, useQueryClient } from '@tanstack/react-query';
import { userApi } from '../api/userApi';
import { queryKeys } from '../../../lib/api/queryClient';
import type { UpdateCertificationRequest, Certification } from '@skillsier/types';

interface UpdateCertificationVariables {
  certificationId: string;
  data: UpdateCertificationRequest;
}

export function useUpdateCertification() {
  const queryClient = useQueryClient();

  return useMutation<Certification, Error, UpdateCertificationVariables>({
    mutationFn: ({ certificationId, data }) => userApi.updateCertification(certificationId, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKeys.users.certifications });
      queryClient.invalidateQueries({ queryKey: queryKeys.users.freelancerProfile });
    },
  });
}