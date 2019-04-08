import React, { PureComponent } from 'react';
import FooterView from './Views';
import {
  Animated,
  Keyboard,
  EmitterSubscription,
  KeyboardEvent,
} from 'react-native';
import { Scaled } from '../../../themes';

const { height } = Scaled.screen;
const defaultKeyboardPadding = height * (Scaled.isXGen ? 0.02 : 0); // Accounts for X-Gen iPhones
let keyboardPadding = new Animated.Value(defaultKeyboardPadding);
let previousLineHeight: number = 0;

export interface Props {}

interface State {
  renderSendArrow: boolean;
  message: string;
}

class Footer extends PureComponent<Props, State> {
  constructor(props: Props) {
    super(props);

    this.state = {
      renderSendArrow: false,
      message: '',
    };

    this.onMessageChange = this.onMessageChange.bind(this);
  }

  keyboardWillShowListener!: EmitterSubscription;
  keyboardWillHideListener!: EmitterSubscription;

  componentDidMount() {
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
    Animated.timing(keyboardPadding, {
      toValue: keyboardHeight,
      duration: 250,
    }).start();
  }

  keyboardWillHide() {
    Animated.timing(keyboardPadding, {
      toValue: defaultKeyboardPadding,
      duration: 250,
    }).start();
  }

  onMessageChange(newMessage: string) {
    this.setState({ message: newMessage });
  }


  render() {
    return (
      <FooterView
        message={this.state.message}
        onMessageChange={this.onMessageChange}
        keyboardPadding={keyboardPadding}
        renderSendArrow={!!this.state.message}
      />
    );
  }
}

export default Footer;
