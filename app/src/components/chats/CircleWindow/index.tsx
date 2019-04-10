import React, { PureComponent } from 'react';
import { Animated } from 'react-native';
import { Circle } from '../ScrollList';
import CircleWindowView from './Views';

export interface Props {
  circles: Circle[];
  circleWindowHeight: Animated.Value;
  onPressCircle(): void;
}

interface State {}

class CircleWindow extends PureComponent<Props, State> {
  render() {
    return <CircleWindowView {...this.props} />;
  }
}

export default CircleWindow;
