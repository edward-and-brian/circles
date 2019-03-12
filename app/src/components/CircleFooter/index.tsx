import React, { PureComponent } from 'react';
import CircleFooterView from './Views';
import { Animated } from 'react-native';

export interface Props {
  height: Animated.Value;
  message: string;
  onMessageChange(newMessage: string): void;
}

interface State {
  renderSendArrow: boolean;
}

class CircleFooter extends PureComponent<Props, State> {
  constructor(props: Props) {
    super(props);

    this.state = {
      renderSendArrow: false,
    };
  }

  static getDerivedStateFromProps(nextProps: Props, prevState: State) {
    if (!!nextProps.message !== prevState.renderSendArrow) {
      console.log('flop');
      return { renderSendArrow: !prevState.renderSendArrow };
    }
    return null;
  }

  render() {
    return (
      <CircleFooterView
        {...this.props}
        renderSendArrow={this.state.renderSendArrow}
      />
    );
  }
}

export default CircleFooter;
