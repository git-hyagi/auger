package encoding

import (
	appsv1 "github.com/openshift/api/apps/v1"
	buildv1 "github.com/openshift/api/build/v1"
	imagev1 "github.com/openshift/api/image/v1"
	oauthv1 "github.com/openshift/api/oauth/v1"
	routev1 "github.com/openshift/api/route/v1"
	templatev1 "github.com/openshift/api/template/v1"
	userv1 "github.com/openshift/api/user/v1"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var (
	ocpScheme = runtime.NewScheme()
	ocpCodecs = serializer.NewCodecFactory(ocpScheme)
)

func tryCodecOCPObject(encoder runtime.Encoder, decoder runtime.Decoder, gv schema.GroupVersion) runtime.Codec {
	switch gv.String() {
	case "apps.openshift.io/v1":
		if err := appsv1.AddToScheme(ocpScheme); err != nil {
			panic(err)
		}
		return ocpCodecs.CodecForVersions(encoder, decoder, appsv1.GroupVersion, appsv1.SchemeGroupVersion)
	case "build.openshift.io/v1":
		if err := buildv1.AddToScheme(ocpScheme); err != nil {
			panic(err)
		}
		return ocpCodecs.CodecForVersions(encoder, decoder, buildv1.GroupVersion, buildv1.SchemeGroupVersion)
	case "image.openshift.io/v1":
		if err := imagev1.AddToScheme(ocpScheme); err != nil {
			panic(err)
		}
		return ocpCodecs.CodecForVersions(encoder, decoder, imagev1.GroupVersion, imagev1.SchemeGroupVersion)
	case "oauth.openshift.io/v1":
		if err := oauthv1.AddToScheme(ocpScheme); err != nil {
			panic(err)
		}
		return ocpCodecs.CodecForVersions(encoder, decoder, oauthv1.GroupVersion, oauthv1.SchemeGroupVersion)
	case "route.openshift.io/v1":
		if err := routev1.AddToScheme(ocpScheme); err != nil {
			panic(err)
		}
		return ocpCodecs.CodecForVersions(encoder, decoder, routev1.GroupVersion, routev1.SchemeGroupVersion)
	case "template.openshift.io/v1":
		if err := templatev1.AddToScheme(ocpScheme); err != nil {
			panic(err)
		}
		return ocpCodecs.CodecForVersions(encoder, decoder, templatev1.GroupVersion, templatev1.SchemeGroupVersion)
	case "user.openshift.io/v1":
		if err := userv1.AddToScheme(ocpScheme); err != nil {
			panic(err)
		}
		return ocpCodecs.CodecForVersions(encoder, decoder, userv1.GroupVersion, userv1.SchemeGroupVersion)
	}
	return nil
}

func tryDecodeOCPObject(gv schema.GroupVersion) runtime.Decoder {
	switch gv.String() {
	case "apps.openshift.io/v1":
		if err := appsv1.AddToScheme(ocpScheme); err != nil {
			panic(err)
		}
		return ocpCodecs.UniversalDecoder(appsv1.SchemeGroupVersion)
	case "build.openshift.io/v1":
		if err := buildv1.AddToScheme(ocpScheme); err != nil {
			panic(err)
		}
		return ocpCodecs.UniversalDecoder(buildv1.SchemeGroupVersion)
	case "image.openshift.io/v1":
		if err := imagev1.AddToScheme(ocpScheme); err != nil {
			panic(err)
		}
		return ocpCodecs.UniversalDecoder(imagev1.SchemeGroupVersion)
	case "oauth.openshift.io/v1":
		if err := oauthv1.AddToScheme(ocpScheme); err != nil {
			panic(err)
		}
		return ocpCodecs.UniversalDecoder(oauthv1.SchemeGroupVersion)
	case "route.openshift.io/v1":
		if err := routev1.AddToScheme(ocpScheme); err != nil {
			panic(err)
		}
		return ocpCodecs.UniversalDecoder(routev1.SchemeGroupVersion)
	case "template.openshift.io/v1":
		if err := templatev1.AddToScheme(ocpScheme); err != nil {
			panic(err)
		}
		return ocpCodecs.UniversalDecoder(templatev1.SchemeGroupVersion)
	case "user.openshift.io/v1":
		if err := userv1.AddToScheme(ocpScheme); err != nil {
			panic(err)
		}
		return ocpCodecs.UniversalDecoder(userv1.SchemeGroupVersion)
	}
	return nil
}
