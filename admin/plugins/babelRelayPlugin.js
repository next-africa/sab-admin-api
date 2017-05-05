/**
 * Created by pdiouf on 2017-05-04.
 */

// `babel-relay-plugin` returns a function for creating plugin instances
const getBabelRelayPlugin = require('babel-relay-plugin');

// load previously saved schema data (see "Schema JSON" below)
const schemaData = require('../schema/schema.json');


// compile code with babel using the plugin
module.exports = getBabelRelayPlugin(schemaData.data);

