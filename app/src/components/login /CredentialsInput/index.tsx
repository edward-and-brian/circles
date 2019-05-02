import React, { PureComponent } from 'react';
import CredentialsInputView from './Views';

export interface Props {
  value: string;
  onChangeText(newValue: string): void;
  placeholder: string;
  label: string;
}

interface State {}

class CredentialsInput extends PureComponent<Props, State> {
  render() {
    return <CredentialsInputView {...this.props} />;
  }
}

export default CredentialsInput;
