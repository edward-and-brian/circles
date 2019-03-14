import React, { PureComponent } from 'react';
import CircleMessageDividerView from './Views';

export interface Props {
  label: string;
}
interface State {}

class CircleMessageDivider extends PureComponent<Props, State> {
  render() {
    return <CircleMessageDividerView {...this.props} />;
  }
}

export default CircleMessageDivider;
