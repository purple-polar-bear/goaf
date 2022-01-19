package geopackage

import (
	"log"
	"net/http"
	"oaf-server/core"
	"github.com/purple-polar-bear/go-ogc-api/features/models"
	"github.com/purple-polar-bear/go-ogc-api/features/services"
	"github.com/purple-polar-bear/go-ogc-api/core/services"

	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/encoding/geojson"
	"github.com/imdario/mergo"
)

type featureService struct {
	geopackage *GeoPackage
	config     *core.Config
}

func Init(config core.Config) featureservices.FeatureService {
	f := &featureService{}

	gp, err := NewGeoPackage(config.Datasource.Geopackage.File, config.Datasource.Geopackage.Fid)
	f.geopackage = &gp
	f.config = &config

	collections := config.Datasource.Collections
	updatedcollections := core.Collections{}

	if len(config.Datasource.Collections) != 0 {
		for _, gpkgc := range f.geopackage.Collections {
			for _, configc := range collections {
				if gpkgc.Identifier == configc.Identifier {
					err = mergo.Merge(&configc, gpkgc)

					if err != nil {
						log.Fatalln(err)
					}
				}
				if (configc.BBox[0] == configc.BBox[2]) == (configc.BBox[1] == configc.BBox[3]) {
					configc.BBox = gpkgc.BBox
				}

				updatedcollections = append(updatedcollections, configc)
			}
		}
		config.Datasource.Collections = updatedcollections
	} else {
		config.Datasource.Collections = f.geopackage.Collections
	}

	if err != nil {
		log.Fatal("Server initialisation error: ", err)
	}

	return f
}

func (service *featureService) Collections() []featuremodels.Collection {

	collections := []featuremodels.Collection{}

	for _, c := range service.geopackage.Collections {
		col := &collection{id: c.Tablename, title: c.Identifier, description: c.Description}
		collections = append(collections, col)
	}

	return collections
}

type collection struct {
	id          string
	title       string
	description string
}

func (c *collection) Id() string         { return c.id }
func (c *collection) Title() string      { return c.title }
func (c collection) Description() string { return c.description }

func (service *featureService) Collection(name string) featuremodels.Collection {

	for _, c := range service.geopackage.Collections {
		if c.Identifier == name {
			return &collection{id: c.Tablename, title: c.Identifier, description: c.Description}
		}
	}

	return nil
}

func (service *featureService) Features(r *http.Request, params *featuremodels.FeaturesParams) featuremodels.Features {
	cn, err := service.config.Datasource.Collections.GetCollections(params.CollectionId)
	if err != nil {
		log.Fatal(err)
	}

	collectionId := params.CollectionId
	offsetParam := uint64(params.Offset)
	limitParam := uint64(params.Limit)
	bboxParam := params.Bbox

	fcGeoJSON, err := service.geopackage.GetFeatures(r.Context(), service.geopackage.DB, cn, params, collectionId, offsetParam, limitParam, nil, bboxParam)

	if err != nil {
		log.Fatal("FeatureColletion error: ", err)
	}

	return fcGeoJSON
}

func (service *featureService) Feature(collectionId string, id string) *featuremodels.Feature {
	geometry := geom.Point{4.873270473933632, 53.083485031473046}
	properties := make(map[string]interface{})
	properties["component_addressareaname"] = "Oosterend"

	return &featuremodels.Feature{
		Feature: core.Feature{
			ID: id,
			Feature: geojson.Feature{
				Geometry:   geojson.Geometry{Geometry: geometry},
				Properties: properties,
			},
		},
	}
}

func (service *featureService) BuildOpenAPISpecification(builder coreservices.OpenAPIBuilder) {
	builderService := featureservices.NewFeatureServiceOpenAPIBuilder(builder)
	builderService.AddBBox()
	builderService.AddDatetime()
  builderService.AddLimit()
  builderService.AddOffset()
}
