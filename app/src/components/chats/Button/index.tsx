import React, { PureComponent } from 'react';
import { Animated } from 'react-native';
import { Chat } from '../ScrollList';
import { Scaled } from '../../../themes';
import ButtonView from './Views';

const openCircleWindowHeight = Scaled.screen.height * 0.25;

export interface Props {
  chat: Chat;
  onPressChat(): void;
}

interface State {}

class Button extends PureComponent<Props, State> {
  constructor(props: Props) {
    super(props);
    this.showCircleWindow = false;
    this.circleWindowHeight = new Animated.Value(0);
    this.toggleCircleWindow = this.toggleCircleWindow.bind(this);
  }

  showCircleWindow: boolean;
  circleWindowHeight: Animated.Value;

  toggleCircleWindow() {
    const newHeight = this.showCircleWindow ? 0 : openCircleWindowHeight;
    this.showCircleWindow = !this.showCircleWindow;

    Animated.timing(this.circleWindowHeight, {
      toValue: newHeight,
      duration: 300,
    }).start();
  }

  render() {
    return (
      <ButtonView
        {...this.props}
        circleWindowHeight={this.circleWindowHeight}
        toggleCircleWindow={this.toggleCircleWindow}
      />
    );
  }
}

export default Button;
