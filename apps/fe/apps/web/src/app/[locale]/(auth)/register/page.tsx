'use client';

import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { useState } from 'react';
import { Button, Input, Card } from '@skillsier/ui';
import { useRegister } from '@skillsier/shared';
import { Mail, Lock, User, AlertCircle, CheckCircle2 } from 'lucide-react';

export default function RegisterPage() {
  const router = useRouter();
  const { mutate: register, isPending, error } = useRegister();
  const [formData, setFormData] = useState({
    email: '',
    username: '',
    password: '',
    confirmPassword: '',
    firstName: '',
    lastName: '',
  });
  const [passwordValidation, setPasswordValidation] = useState({
    length: false,
    uppercase: false,
    lowercase: false,
    number: false,
  });

  const validatePassword = (password: string) => {
    setPasswordValidation({
      length: password.length >= 8,
      uppercase: /[A-Z]/.test(password),
      lowercase: /[a-z]/.test(password),
      number: /[0-9]/.test(password),
    });
  };

  const handlePasswordChange = (password: string) => {
    setFormData({ ...formData, password });
    validatePassword(password);
  };

  const isPasswordValid = Object.values(passwordValidation).every(Boolean);
  const passwordsMatch = formData.password === formData.confirmPassword;

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!isPasswordValid || !passwordsMatch) return;

    const { confirmPassword, ...registerData } = formData;
    register(registerData, {
      onSuccess: () => {
        router.push('/dashboard');
      },
    });
  };

  return (
    <Card padding="lg" shadow="lg">
      <div className="text-center mb-8">
        <Link href="/" className="inline-flex items-center gap-2 mb-4">
          <div