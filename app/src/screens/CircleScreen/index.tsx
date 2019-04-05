import React, { PureComponent } from 'react';
import {
  Animated,
  Keyboard,
  EmitterSubscription,
  KeyboardEvent,
} from 'react-native';
import NavigationService from '../../navigation/NavigationService';
import CircleScreenView from './Views';

export interface Props {}
interface State {
  message: string;
}

class CircleScreen extends PureComponent<Props, State> {
  onPressBack() {
    NavigationService.back();
  }

  render() {
    return (
      <CircleScreenView
        onPressBack={this.onPressBack}
      />
    );
  }
}

export default CircleScreen;
