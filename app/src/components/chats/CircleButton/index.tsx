import React, { PureComponent } from 'react';
import { Circle } from '../ScrollList';
import CircleButtonView from './Views';

export interface Props {
  circle: Circle;
  onPressCircle(): void;
}

interface State {}

class CircleButton extends PureComponent<Props, State> {
  render() {
    return <CircleButtonView {...this.props} />;
  }
}

export default CircleButton;
