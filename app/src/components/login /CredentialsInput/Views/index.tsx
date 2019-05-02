import React, { PureComponent } from 'react';
import { Text, View, TextInput } from 'react-native';
import { Colors } from '../../../../themes';
import styles from './styles';

export interface Props {
  value: string;
  onChangeText(newValue: string): void;
  placeholder: string;
  label: string;
}

class CredentialsInputView extends PureComponent<Props> {
  render() {
    return (
      <View style={styles.container}>
        <Text style={styles.label}>{this.props.label}</Text>
        <TextInput
          style={styles.input}
          value={this.props.value}
          onChangeText={this.props.onChangeText}
          placeholder={this.props.placeholder}
          placeholderTextColor={Colors.darkGray}
          autoCapitalize="none"
          autoCorrect={false}
        />
      </View>
    );
  }
}

export default CredentialsInputView;
