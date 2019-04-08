import React, { PureComponent } from 'react';
import { Animated } from 'react-native';
import { Chat } from '../ScrollList';
import { Scaled } from '../../../themes';
import ButtonView from './Views';

const openCircleWindowHeight = Scaled.screen.height * 0.25;
let showCircleWindow = false;
let circleWindowHeight = new Animated.Value(0);

export interface Props {
  chat: Chat;
  onPressChat(): void;
}

interface State {}

class Button extends PureComponent<Props, State> {
  constructor(props: Props) {
    super(props);

    this.toggleCircleWindow = this.toggleCircleWindow.bind(this);
  }

  toggleCircleWindow() {
    const prevValue = showCircleWindow;
    showCircleWindow = !prevValue;

    const newHeight = prevValue ? 0 : openCircleWindowHeight;
    Animated.timing(circleWindowHeight, {
      toValue: newHeight,
    }).start();
  }

  render() {
    return (
      <ButtonView
        {...this.props}
        circleWindowHeight={circleWindowHeight}
        toggleCircleWindow={this.toggleCircleWindow}
      />
    );
  }
}

export default Button;
