import React, { PureComponent } from 'react';
import NavbarView from './Views';
import { StyleProp, ViewStyle } from 'react-native';

export interface Props {
  onPressBack(): void;
  style?: StyleProp<ViewStyle>;
}

class Navbar extends PureComponent<Props> {
  render() {
    return <NavbarView {...this.props} />;
  }
}

export default Navbar;
