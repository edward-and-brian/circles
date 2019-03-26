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
const textInputHeight = height * (Scaled.isXGen ? 0.053 : 0.065);
const heightDefault = textInputHeight + height * (Scaled.isXGen ? 0.02 : 0);
let footerHeight = new Animated.Value(heightDefault);

export interface Props {
  message: string;
  onMessageChange(newMessage: string): void;
}

interface State {
  renderSendArrow: boolean;
}

class Footer extends PureComponent<Props, State> {
  constructor(props: Props) {
    super(props);

    this.state = {
      renderSendArrow: false,
    };
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

  static getDerivedStateFromProps(nextProps: Props, prevState: State) {
    if (!!nextProps.message !== prevState.renderSendArrow) {
      return { renderSendArrow: !prevState.renderSendArrow };
    }
    return null;
  }

  render() {
    return (
      <FooterView
        {...this.props}
        height={footerHeight}
        renderSendArrow={this.state.renderSendArrow}
      />
    );
  }
}

export default Footer;
