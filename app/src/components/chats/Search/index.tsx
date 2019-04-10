import React, { PureComponent } from 'react';
import SearchView from './Views';

export interface Props {}

class Search extends PureComponent<Props> {
  render() {
    return <SearchView {...this.props} />;
  }
}

export default Search;
