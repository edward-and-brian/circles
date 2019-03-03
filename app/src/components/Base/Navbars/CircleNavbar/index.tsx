import React, { PureComponent } from 'react';
import CircleNavbarView from './Views';

export interface Props {
  onPressBack(): void;
}

class CircleNavbar extends PureComponent<Props> {
  render() {
    return <CircleNavbarView {...this.props} />;
  }
}

export default CircleNavbar;
