/**
 * Components that are present to check that navigation works.
 * They must be removed when components will be provided by other
 * plugins.
 *
 * FIXME Removes these components.
 */
'use strict';

import React from 'react';

const Plugin1Feature1 = () => (<div>Plugin1 feature 1</div>);
const Plugin2Feature1 = () => (<div>Plugin2 feature 1</div>);
const Plugin2Feature2 = () => (<div>Plugin2 feature 2</div>);

export {Plugin1Feature1, Plugin2Feature1, Plugin2Feature2};