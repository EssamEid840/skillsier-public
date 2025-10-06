import React from "react";
import { useLocalSearchParams } from "expo-router";
import Screen from "../../../../src/navigation/Screen";
export default function VerificationStep(){
  const { step } = useLocalSearchParams<{ step: string }>();
  return <Screen title={`Verification · ${step ?? ""}`} />;
}
