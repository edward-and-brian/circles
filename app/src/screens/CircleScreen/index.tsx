import React, { PureComponent } from 'react';
import {
  Animated,
  Keyboard,
  EmitterSubscription,
  KeyboardEvent,
} from 'react-native';
import NavigationService from '../../navigation/NavigationService';
import CircleScreenView from './Views';
import { Scaled } from '../../themes';

const textInputHeight = Scaled.screen.height * (Scaled.isXGen ? 0.053 : 0.065);
const heightDefault =
  textInputHeight + Scaled.screen.height * (Scaled.isXGen ? 0.02 : 0);
let footerHeight = new Animated.Value(heightDefault);

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

  keyboardWillShowListener!: EmitterSubscription;
  keyboardWillHideListener!: EmitterSubscription;

  componentWillMount() {
    this.keyboardWillShowListener = Keyboard.addListener(
      'keyboardWillShow',
      this.keyboardWillShow,
    );
    this.keyboardWillHideListener = Keyboard.addListener(
      'keyboardWillHide',
      this.keyboardWillHide,
    );
  }

  componentWillUnmount() {
    this.keyboardWillShowListener.remove();
    this.keyboardWillHideListener.remove();
  }

  keyboardWillShow(event: KeyboardEvent) {
    const keyboardHeight = event.endCoordinates.height;
    Animated.timing(footerHeight, {
      toValue: keyboardHeight + textInputHeight,
      duration: 250,
    }).start();
  }

  keyboardWillHide() {
    Animated.timing(footerHeight, {
      toValue: heightDefault,
      duration: 250,
    }).start();
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
        footerHeight={footerHeight}
        message={this.state.message}
        onMessageChange={this.onMessageChange}
      />
    );
  }
}

export default CircleScreen;
