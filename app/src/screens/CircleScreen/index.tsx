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

const textInputHeight = Scaled.screen.height * 0.055;
const heightDefault =
  textInputHeight + (Scaled.isXGen ? Scaled.screen.height * 0.025 : 0);
let footerHeight = new Animated.Value(heightDefault);

export interface Props {}

class CircleScreen extends PureComponent<Props> {
  constructor(props: Props) {
    super(props);
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

  render() {
    return (
      <CircleScreenView
        onPressBack={this.onPressBack}
        footerHeight={footerHeight}
      />
    );
  }
}

export default CircleScreen;
