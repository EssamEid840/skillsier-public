import { describe, it, expect } from 'vitest';
import { SUPPORTED_LOCALES, isRTL } from '../src/config';

describe('i18n config', () => {
  it('supports en and ar locales', () => {
    expect(SUPPORTED_LOCALES).toEqual(['en', 'ar']);
  });

  it('identifies RTL locales correctly', () => {
    expect(isRTL('ar')).toBe(true);
    expect(isRTL('en')).toBe(false);
  });
});