import React, { PureComponent } from 'react';
import CircleAvatarButtonView from './Views';

class CircleAvatarButton extends PureComponent {
  render() {
    return <CircleAvatarButtonView {...this.props} />;
  }
}

export default CircleAvatarButton;
