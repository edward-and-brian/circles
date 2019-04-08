import React, { PureComponent } from 'react';
import { TouchableOpacity, View } from 'react-native';
import styles from './styles';

export interface Props {
  color: string;
  diameter: number;
  clickable: boolean;
  onPress(): void;
}

class AvatarButtonView extends PureComponent<Props> {
  render() {
    const { color, diameter } = this.props;
    return this.props.clickable ? (
      <TouchableOpacity
        style={{
          backgroundColor: color,
          width: diameter,
          height: diameter,
          borderRadius: diameter / 2,
        }}
      />
    ) : (
      <View
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

export default AvatarButtonView;
