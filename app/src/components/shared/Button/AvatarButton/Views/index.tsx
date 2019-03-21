import React, { PureComponent } from 'react';
import { TouchableOpacity } from 'react-native';
import styles from './styles';

export interface Props {
  color: string;
  diameter: number;
}

class CircleAvatarButtonView extends PureComponent<Props> {
  render() {
    const { color, diameter } = this.props;
    return (
      <TouchableOpacity
        style={{
          backgroundColor: color,
          width: diameter,
          height: diameter,
          borderRadius: diameter / 2,
        }}
      />
    );
  }
}

export default CircleAvatarButtonView;
