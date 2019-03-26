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
  constructor(props: Props) {
    super(props);

    this.state = {
      message: '',
    };

    this.onMessageChange = this.onMessageChange.bind(this);
  }

  onPressBack() {
    NavigationService.back();
  }

  onMessageChange(newMessage: string) {
    this.setState({ message: newMessage });
  }

  render() {
    return (
      <CircleScreenView
        onPressBack={this.onPressBack}
        message={this.state.message}
        onMessageChange={this.onMessageChange}
      />
    );
  }
}

export default CircleScreen;
