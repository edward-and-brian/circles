import React, { PureComponent } from 'react';
import AvatarButtonView from './Views';
import { Colors } from '../../../../themes';

export interface Props {
  color: string;
  diameter: number;
  clickable: boolean;
  onPress(): void;
}

class AvatarButton extends PureComponent<Props> {
  static defaultProps: Partial<Props> = {
    color: Colors.yellow,
    clickable: false,
    onPress: () => {},
  };

  render() {
    return <AvatarButtonView {...this.props} />;
  }
}

export default AvatarButton;
