import React, { PureComponent } from 'react';
import NavigationService from '../../navigation/NavigationService';
import SignInScreenView from './Views';

export interface Props {}
interface State {}

class SignInScreen extends PureComponent<Props, State> {
  onPressBack() {
    NavigationService.back();
  }

  render() {
    return <SignInScreenView />;
  }
}

export default SignInScreen;
