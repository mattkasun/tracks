#include <gtk/gtk.h>
#include <osm-gps-map.h>


GtkWidget *map, *radio1, *radio2, *radio3;

int main (int argc, char **argv)
{
  GtkBuilder *builder;
  GtkWidget *window, *map_box, *button;
  GError *error = NULL;
  //OsmGpsMap *map;
  //OsmGpsMapLayer *osd;

  gtk_init (&argc, &argv);

  OsmGpsMapSource_t source = OSM_GPS_MAP_SOURCE_GOOGLE_SATELLITE;

    if ( !osm_gps_map_source_is_valid(source) )
        return 1;


  builder = gtk_builder_new ();
  gtk_builder_add_from_file(builder, "map.ui", &error);
  if (error)
  {
    g_error ("ERROR: %s\n", error->message);
    return 2;
  }
  button = gtk_button_new ();
  window = GTK_WIDGET (gtk_builder_get_object (builder, "window1"));
  map = g_object_new (OSM_TYPE_GPS_MAP,
                     //"map-source", source,
                     //"tile-cache", "/tmp/",
                     "proxy-uri", g_getenv("http_proxy"),
                      NULL);
  GtkWidget *osd = g_object_new (OSM_TYPE_GPS_MAP_OSD,
                        "show-scale",TRUE,
                        "show-coordinates",TRUE,
                        "show-crosshair",TRUE,
                        "show-dpad",FALSE,
                        "show-zoom",TRUE,
                        "show-gps-in-dpad",FALSE,
                        "show-gps-in-zoom",FALSE,
                        "dpad-radius", 30,
                        "show-copyright", TRUE,
                        NULL);

  osm_gps_map_layer_add (OSM_GPS_MAP(map), OSM_GPS_MAP_LAYER(osd));
  map_box = GTK_WIDGET(gtk_builder_get_object(builder, "map_box"));
  radio1 = GTK_WIDGET(gtk_builder_get_object(builder, "radio_openstreet"));
  radio2 = GTK_WIDGET(gtk_builder_get_object(builder, "radio_goo_street"));
  radio3 = GTK_WIDGET(gtk_builder_get_object(builder, "radio_goo_hybrid"));
  if (map_box == NULL) {
    return  3;
  }
  //gtk_container_add(map_box, GTK_WIDGET(map));
  //gtk_container_add(
    //  GTK_CONTAINER(gtk_builder_get_object(builder, "map_box")),
      //GTK_WIDGET(map));

  gtk_box_pack_start(
      GTK_BOX(gtk_builder_get_object(builder, "map_box")),
      GTK_WIDGET(map), TRUE, TRUE, 0);
  gtk_builder_connect_signals(builder, NULL);

  g_object_unref(G_OBJECT(builder));

  gtk_widget_show_all (window);

  gtk_main();

  return 0;
}

//void change_source (GtkToggleButton *button, GtkWidget *radio1, GtkWidget *radio2, GtkWidget *radio3) 
void change_source ()
{
  //GValue *value = g_value_init (value, 1);
  g_print("g_print source changed\n");
  //g_log ("g_log source changed");
  gboolean done;
  printf("source changed\n");
  if (gtk_toggle_button_get_active(GTK_TOGGLE_BUTTON (radio1))) 
  {
    printf("source is %d\n", OSM_GPS_MAP_SOURCE_OPENSTREETMAP);
    //g_value_set_int(value, (OSM_GPS_MAP_SOURCE_OPENSTREETMAP));
    g_object_set(map, "map-source", OSM_GPS_MAP_SOURCE_OPENSTREETMAP, NULL);
  }
  else if (gtk_toggle_button_get_active(GTK_TOGGLE_BUTTON (radio2))) 
  {
   g_object_set(map, "map-source", OSM_GPS_MAP_SOURCE_GOOGLE_STREET, NULL);
  }
  else if (gtk_toggle_button_get_active(GTK_TOGGLE_BUTTON (radio3))) 
  {
    g_object_set (map, "map-source", OSM_GPS_MAP_SOURCE_VIRTUAL_EARTH_SATELLITE, NULL);
  }
  //done = osm_gps_map_map_redraw (OSM_GPS_MAP(map));
  gtk_widget_hide(map);
  gtk_widget_show(map);
}
