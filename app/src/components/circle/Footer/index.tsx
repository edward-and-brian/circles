import React, { PureComponent } from 'react';
import FooterView from './Views';
import { Animated } from 'react-native';

export interface Props {
  height: Animated.Value;
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
        renderSendArrow={this.state.renderSendArrow}
      />
    );
  }
}

export default Footer;
