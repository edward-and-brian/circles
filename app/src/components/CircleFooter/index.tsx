import React, { PureComponent } from 'react';
import CircleFooterView from './Views';
import { Animated } from 'react-native';

export interface Props {
  height: Animated.Value;
}

interface State {
  message: string;
}

class CircleFooter extends PureComponent<Props, State> {
  constructor(props: Props) {
    super(props);

    this.state = {
      message: '',
    };
  }

  render() {
    return <CircleFooterView {...this.props} message={this.state.message} />;
  }
}

export default CircleFooter;
