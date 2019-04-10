import React, { PureComponent } from 'react';
import SearchView from './Views';

export interface Props {}
interface State {
  message: string;
}

class Search extends PureComponent<Props, State> {
  constructor(props: Props) {
    super(props);

    this.state = {
      message: '',
    };

    this.onMessageChange = this.onMessageChange.bind(this);
  }

  onMessageChange(newMessage: string) {
    this.setState({ message: newMessage });
  }

  render() {
    return (
      <SearchView
        {...this.props}
        message={this.state.message}
        onMessageChange={this.onMessageChange}
      />
    );
  }
}

export default Search;
