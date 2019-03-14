import React, { PureComponent } from 'react';
import CircleAvatarButtonView from './Views';
import { Colors, Scaled } from '../../../../themes';

export interface Props {
  color: string;
  diameter: number;
}

class CircleAvatarButton extends PureComponent<Props> {
  static defaultProps: Partial<Props> = {
    color: Colors.yellow,
    diameter: Scaled.screen.height * 0.055,
  };

  render() {
    return <CircleAvatarButtonView {...this.props} />;
  }
}

export default CircleAvatarButton;
