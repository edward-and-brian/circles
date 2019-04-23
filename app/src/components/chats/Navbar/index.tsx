import React, { PureComponent } from 'react';
import NavbarView from './Views';

export interface Props {}

class Navbar extends PureComponent<Props> {
  render() {
    return <NavbarView {...this.props} />;
  }
}

export default Navbar;
