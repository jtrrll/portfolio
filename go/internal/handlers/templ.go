package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	meter         = otel.Meter("portfolio")
	tracer        = otel.Tracer("portfolio")
	pageViewCount metric.Int64Counter
)

func init() {
	var err error
	pageViewCount, err = meter.Int64Counter(
		"page.views",
		metric.WithDescription("Number of page views"),
		metric.WithUnit("{view}"),
	)
	if err != nil {
		panic(err)
	}
}

func TemplPage(component templ.Component) echo.HandlerFunc {
	templHandler := echo.WrapHandler(templ.Handler(component, templ.WithStreaming()))
	return func(c echo.Context) error {
		ctx, span := tracer.Start(c.Request().Context(), "page.render",
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()

		c.SetRequest(c.Request().WithContext(ctx))

		pageViewCount.Add(c.Request().Context(), 1,
			metric.WithAttributeSet(attribute.NewSet(
				semconv.HTTPRoute(c.Path()),
			)),
		)

		return templHandler(c)
	}
}
