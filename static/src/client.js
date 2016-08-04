
var Swagger = require('swagger-client');

//var client = new Swagger({
//    url: '/swagger.json',
//    success: function() {
//        console.log(client);
        //client.notes.get_notes_date_from_date_to({
        //    date_from: '2222-22-22',
        //    date_to: '2222-22-22'
        //}, {responseContentType: 'application/json'}, function(pet){
        //    console.log('pet', pet.data);
        //});

        //client.notes.post_notes({name: "okok"}, function(pet){
        //    console.log('pet', pet);
        //});

        //client.notes.help();
        //client.inflow.help();
    //}
//});

var $ = require("jquery");
var client = new $.RestClient('/api/');
client.add("notes");
client.notes.read(2);