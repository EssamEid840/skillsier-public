import * as React from 'react';
import {
  TextInput,
  Text,
  View,
  StyleSheet,
  type TextInputProps,
} from 'react-native';
import { colors } from '../../tokens';

export interface InputProps extends TextInputProps {
  error?: string;
  label?: string;
}

export const Input = React.forwardRef<TextInput, InputProps>(
  ({ error, label, style, ...props }, ref) => {
    return (
      <View style={styles.container}>
        {label && <Text style={styles.label}>{label}</Text>}
        <TextInput
          ref={ref}
          style={[
            styles.input,
            error && styles.inputError,
            props.editable === false && styles.disabled,
            style,
          ]}
          placeholderTextColor={colors.secondary[400]}
          {...props}
        />
        {error && <Text style={styles.error}>{error}</Text>}
      </View>
    );
  }
);

Input.displayName = 'Input';

const styles = StyleSheet.create({
  container: {
    width: '100%',
  },
  label: {
    marginBottom: 8,
    fontSize: 14,
    fontWeight: '500',
    color: colors.secondary[900],
  },
  input: {
    height: 40,
    width: '100%',
    borderRadius: 8,
    borderWidth: 1,
    borderColor: colors.secondary[300],
    backgroundColor: '#FFFFFF',
    paddingHorizontal: 12,
    paddingVertical: 8,
    fontSize: 14,
    color: colors.secondary[900],
  },
  inputError: {
    borderColor: colors.error.DEFAULT,
  },
  disabled: {
    opacity: 0.5,
  },
  error: {
    marginTop: 4,
    fontSize: 12,
    color: colors.error.DEFAULT,
  },
});