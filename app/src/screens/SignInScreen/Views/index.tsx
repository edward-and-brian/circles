import React, { PureComponent } from 'react';
import { View, Text } from 'react-native';
import NavigationService from '../../../navigation/NavigationService';
import CredentialsInput from '../../../components/login /CredentialsInput';
import LoginButton from '../../../components/shared/Button/LoginButton';
import { Colors } from '../../../themes';
import styles from './styles';

export interface Props {
  username: string;
  phone: string;
  onChangeUsername(newUsername: string): void;
  onChangePhone(newPhone: string): void;
  login(): void;
}

class SignInScreenView extends PureComponent<Props> {
  render() {
    return (
      <View style={styles.container}>
        <View style={styles.titleContainer}>
          <Text style={styles.title}>Circles</Text>
        </View>
        <View style={styles.inputsContainer}>
          <CredentialsInput
            value={this.props.username}
            onChangeText={this.props.onChangeUsername}
            placeholder="edlo"
            label="username"
          />
          <CredentialsInput
            value={this.props.phone}
            onChangeText={this.props.onChangePhone}
            placeholder="626.607.7099"
            label="phone"
          />
        </View>
        <View style={styles.buttonsContainer}>
          <View style={styles.buttonContainer}>
            <LoginButton
              text="Login"
              color={Colors.c_blue}
              onPress={this.props.login}
            />
          </View>
          <View style={styles.buttonContainer}>
            <LoginButton
              text="New to Circles"
              color={Colors.darkGray}
              fontColor={Colors.white}
              onPress={() => NavigationService.navigate('chats')}
            />
          </View>
        </View>
      </View>
    );
  }
}

export default SignInScreenView;
