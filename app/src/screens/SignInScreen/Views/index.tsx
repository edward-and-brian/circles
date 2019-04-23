import React, { PureComponent } from 'react';
import { View, Text } from 'react-native';
import styles from './styles';

export interface Props {}

class SignInScreenView extends PureComponent<Props> {
  render() {
    return (
      <View style={styles.container}>
        <View style={styles.titleContainer}>
          <Text style={styles.title}>Circles</Text>
        </View>
        <View style={styles.inputContainer} />
        <View style={styles.buttonContainer} />
      </View>
    );
  }
}

export default SignInScreenView;
